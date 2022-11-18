package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Struct object to gather Data
type ShopifyReview struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Rating   uint8              `json:"rating" bson:"rating"`
	Comment  string             `json:"comment" bson:"comment"`
	Stars    string             `json:"starsOfStars" bson:"starsOfStars"`
	Date     string             `json:"date" bson:"date"`
	ShopName string             `json:"shopName" bson:"shopName"`
}

type ThreeWordPhrases struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Rating   uint8              `json:"rating" bson:"rating"`
	Comment  string             `json:"comment" bson:"comment"`
	Stars    string             `json:"starsOfStars" bson:"starsOfStars"`
	Date     string             `json:"date" bson:"date"`
	ShopName string             `json:"shopName" bson:"shopName"`
}
