package models

import (

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transactions struct {
	ID                primitive.ObjectID `json:"_id" bson:"_id"`
	Account_id        int                `json:"account_id" bson:"account_id"`
	Transaction_count int                `json:"transaction_count" bson:"transaction_count"`
	Bucket_start_date primitive.DateTime `json:"bucket_start_date" bson:"bucket_start_date"`
	Bucket_end_date   primitive.DateTime `json:"bucket_end_date" bson:"bucket_end_date"`
	Transactions      []innerTransaction
}
type innerTransaction struct {
	Date             primitive.DateTime `json:"date" bson:"date"`
	Amount           int                `json:"amount" bson:"amount"`
	Transaction_code string             `json:"transaction_code" bson:"transaction_code"`
	Symbol           string             `json:"symbol" bson:"symbol"`
	Price            string             `json:"price" bson:"price"`
	Total            string             `json:"total" bson:"total"`
}
