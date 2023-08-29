package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	Address     Address            `bson:"address" json:"address"`
	Borough     string             `bson:"borough" json:"borough"`
	Cuisine     string             `bson:"cuisine" json:"cuisine"`
	Grades      []Grade            `bson:"grades" json:"grades"`
	Name        string             `bson:"name" json:"name"`
	RestaurantID string             `bson:"restaurant_id" json:"restaurant_id"`
}

type Address struct {
	Building string    `bson:"building" json:"building"`
	Coord    []float64 `bson:"coord" json:"coord"`
	Street   string    `bson:"street" json:"street"`
	Zipcode  string    `bson:"zipcode" json:"zipcode"`
}

type Grade struct {
	Date  time.Time `bson:"date" json:"date"`
	Grade string    `bson:"grade" json:"grade"`
	Score int       `bson:"score" json:"score"`
}