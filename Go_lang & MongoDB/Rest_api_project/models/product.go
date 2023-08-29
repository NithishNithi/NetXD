package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `Json:"id" bson:"_id"`
	Name        string             `Json:"name" bson:"name,required"`
	Price       float64             `Json:"price" bson:"price,required"`
	Description string             `Json:"description" bson:"description,omitempty"`
}
