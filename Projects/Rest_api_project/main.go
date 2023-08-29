package main

import (
	"fmt"
	"project/service"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
)

func main() {
	fmt.Println("Started Connection")
	// product1, _ := service.Findmodel1()
	// for _, j := range product1 {
	// 	fmt.Println(j.ID,j.Transaction_count)
	// }
	// service.FetchAggregatedTransactions()
	
}
