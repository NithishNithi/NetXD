package service

import (
	"context"
	"fmt"
	"project/config"
	"project/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// func to connect to the database
func Modelcontext() *mongo.Collection {
	client, _ := config.ConnectDatabase()
	return config.Getcollection(client, "sample_analytics", "transactions")
}
func Findmodel1() ([]*models.Transactions, error) {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	// bson M stands for bson map
	// bson D stands for bson Document
	filter := bson.M{}
	result, err := Modelcontext().Find(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		// building an array
		var items []*models.Transactions
		for result.Next(ctx) {
			item := &models.Transactions{}
			err := result.Decode(item)
			if err != nil {
				return nil, err
			}

			items = append(items, item)
		}
		if err := result.Err(); err != nil {
			return nil, err
		}
		if len(items) == 0 {
			return []*models.Transactions{}, nil
		}
		return items, nil

	}

}
func FetchAggregatedTransactions() {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	matchStage := bson.D{{"$match", bson.D{{"transaction_count", 100}}}}
	groupStage := bson.D{
		{
			"$group", bson.D{
				{"_id", "$account_id"},
				{"total_count", bson.D{{"$sum", "$transaction_count"}}},
			}}}
	result, err := Modelcontext().Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		fmt.Println(err.Error)
	} else {
		var showWithInfo []bson.M
		if err = result.All(ctx, &showWithInfo); err != nil {
			panic(err)
		}
		fmt.Println(showWithInfo)
	}
}

func UpdateTransaction(initialValue int, newValue int)(*mongo.UpdateResult,error) {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.D{{"account_id", initialValue}}
	update := bson.D{{"$set",bson.D{{"account_id",newValue}}}}
	result,err:=Modelcontext().UpdateOne(ctx,filter,update)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	return result,nil
}
