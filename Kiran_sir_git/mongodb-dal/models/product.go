package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct{
	ID primitive.ObjectID `json:"id" bson:"_id"`
	Name string           `json:"name" bson:"name,required"`
	Price float64          `json:"price" bson:"price,required"`
	Description string     `json:"description" bson:"description,omitempty"`
}