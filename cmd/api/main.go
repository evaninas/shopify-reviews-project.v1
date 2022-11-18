package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"omnisend-test/config"
	"omnisend-test/models"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Creating slice of a struct
	var allReviews []models.ShopifyReview

	// Getting params from query (skip=, limit=, sortBy=)
	queryValues := r.URL.Query()
	batchSize := int32(50)
	skip, _ := strconv.ParseInt(queryValues.Get("skip"), 10, 32)
	limit, _ := strconv.ParseInt(queryValues.Get("limit"), 0, 32)
	sortBy := queryValues.Get("sortBy")

	// Getting context and db connection from config
	ctx := config.CTX
	db := config.DB

	// Batch sets the number of documents to return in every batch
	// Limit limits results returned
	// Skip specifies number of documents to skip before returning
	options := options.FindOptions{}
	options.BatchSize = &batchSize
	options.Limit = &limit
	options.Skip = &skip

	// Switch statement to sort results by Ascending/Descending order
	switch sortBy {
	case "ratingAscend":
		options.Sort = bson.M{"rating": 1}
	case "ratingDescend":
		options.Sort = bson.M{"rating": -1}
	default:
	}

	// Trying to fetch all reviews from reviewed data, if err - return log
	coll, err := db.Collection("ShopifyReviews").Find(ctx, bson.D{}, &options)
	if err != nil {
		log.Fatal(err)
	}
	defer coll.Close(ctx)

	// Find returns a mongo cursor that can be used to iterate over a collections using (Next)
	for coll.Next(ctx) {
		// Setting element to a bson document address
		var ShopifyReview models.ShopifyReview

		// Decoding document into val
		err = coll.Decode(&ShopifyReview)
		if err != nil {
			log.Fatal("Error at decoding :", err)
		}

		// Pushing single review to all reviews slice
		allReviews = append(allReviews, ShopifyReview)
	}

	// Encoding allReviews to Json format
	reviewsJson, err := json.Marshal(allReviews)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", reviewsJson)
}

func main() {
	r := httprouter.New()
	r.GET("/", Index)

	fmt.Println("Server is listening on port :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
