// Define all models
package models

import (
	// "time"


)

type Transactions struct {
	ID               string `json:"id" bson:"id"`
	Transaction_Date string `json:"transaction_date" bson:"transaction_date"`
	Amount           int                `json:"amount" bson:"amount"`
	Transaction_Type string             `json:"transaction_type" bson:"transaction_type"`
	Customer_Id      string             `json:"customer_id" bson:"customer_id"`
}
type DBResponse struct {
	ID               string `json:"id" bson:"id"`
	Transaction_Date string  `json:"transaction_date" bson:"transaction_date"`
	Amount           int                `json:"amount" bson:"amount"`
	Transaction_Type string             `json:"transaction_type" bson:"transaction_type"`
	Customer_Id      string             `json:"customer_id" bson:"customer_id"`
}
