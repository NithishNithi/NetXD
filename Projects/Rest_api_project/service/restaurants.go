package service

import (
	"context"
	"fmt"
	"project/config"
	"project/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ProductContext1()*mongo.Collection {
	client,_ := config.ConnectDatabase()
	return config.Getcollection(client,"sample_restaurants","restaurants")
}

func FindProducts1() ([]*models.Product, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	filter := bson.D{}
	limit := options.Find().SetLimit(10)
	result, err := ProductContext1().Find(ctx, filter, limit)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		fmt.Println(result)
		var products []*models.Restaurant
		for result.Next(ctx) {
			product := &models.Restaurant{}
			err := result.Decode(product)
			if err != nil {
				return nil, err
			}
			fmt.Println(product)
			products = append(products, product)
		}
		if err := result.Err(); err != nil {
			return nil, err
		}
		if len(products) == 0 {
			return []*models.Product{}, nil
		}
		return []*models.Product{}, nil
	}
}
