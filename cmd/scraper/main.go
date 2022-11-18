package main

import (
	"context"
	"fmt"
	"log"
	"omnisend-test/config"
	"omnisend-test/models"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to mongo database local host, returns err with context
func connectToMongoDB() (client *mongo.Client, err error) {
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:root@localhost:27017")); err != nil {
		return
	}
	err = client.Ping(context.TODO(), nil)
	return
}

// Function that writes reviews to ShopifyReviews collection
func writeToDB(ctx context.Context, db *mongo.Database, review models.ShopifyReview) {
	collection := db.Collection("ShopifyReviews")

	_, err := collection.InsertOne(ctx, bson.D{
		{Key: "rating", Value: review.Rating},
		{Key: "comment", Value: review.Comment},
		{Key: "shopName", Value: review.ShopName},
		{Key: "date", Value: review.Date},
		{Key: "starOfStars", Value: review.Stars},
	})
	if err != nil {
		log.Fatal("Write to DB error: ", err)
	}
	fmt.Print("Inserting review to whole list :", review)
}

// Function that writes reviews to Three-words collection
func writeThreeWords(ctx context.Context, db *mongo.Database, review models.ShopifyReview) {
	collection := db.Collection("ThreeWordPhrases")

	_, err := collection.InsertOne(ctx, bson.D{
		{Key: "rating", Value: review.Rating},
		{Key: "comment", Value: review.Comment},
		{Key: "shopName", Value: review.ShopName},
		{Key: "date", Value: review.Date},
		{Key: "starOfStars", Value: review.Stars},
	})
	if err != nil {
		log.Fatal("Write to DB error: ", err)
	}
	fmt.Print("writing to Three Words collection :", review)
}

func CountWords(s string) int {
	return len(strings.Fields(s))
}

func main() {

	// Getting context and db connection from config
	ctx := config.CTX
	db := config.DB

	// Variable that is linked to main scraping website
	scrapeUrl := "https://apps.shopify.com/omnisend/reviews"

	// Colly parameter
	c := colly.NewCollector(
		colly.MaxDepth(2),
	)

	// Error catcher
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	// Go through HTML and print selected elements
	c.OnHTML("div.review-listing", func(h *colly.HTMLElement) {
		ratingString := h.ChildAttr("div.ui-star-rating", "data-rating")
		shopName := h.ChildText("h3.review-listing-header__text")
		starsOfStars := strings.TrimSpace(h.DOM.Find("div.review-metadata__item .review-metadata__item-value").Last().Text())
		date := h.ChildText("div.review-metadata__item-label")
		comment := h.ChildText("p")

		// Converting string rating to integer rating
		ratingInteger, err := strconv.Atoi(ratingString)
		if err != nil {
			log.Println("Stars str to in err: ", err)
		}
		rating := uint8(ratingInteger)

		// Declaring review with values from scraped reviews
		review := models.ShopifyReview{
			Rating:   rating,
			Date:     date,
			Comment:  comment,
			ShopName: shopName,
			Stars:    starsOfStars,
		}

		writeToDB(ctx, db, review)
		if CountWords(review.Comment) == 3 {
			writeThreeWords(ctx, db, review)
		}

	})

	// Callback function to go through pagination links
	c.OnHTML("a.search-pagination__link", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		// Visit the link
		c.Visit(h.Request.AbsoluteURL(link))
	})

	// Logging which page is being visited
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
		time.Sleep(1 * time.Second)
	})

	c.Visit(scrapeUrl)

	// Returns when the jobs are finished
	c.Wait()
	fmt.Println("Finished")

}
