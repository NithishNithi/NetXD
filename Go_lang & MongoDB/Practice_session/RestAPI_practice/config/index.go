// This file is to define the connection strings
// Port
// Datbase conenctivity
package config

import (
	"context"
	"fmt"
	"Rest-API/constants"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDataBase() (*mongo.Client,error){
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	mongoConnection := options.Client().ApplyURI(constants.ConnectionString)
	mongoClient, err := mongo.Connect(ctx, mongoConnection)
	if err!=nil {
		log.Fatal(err.Error())
		return nil, err
	}
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	fmt.Println("Database Connected")
	return mongoClient,nil	
}

func GetCollection(client *mongo.Client,dbName string,collectionName string) *mongo.Collection{
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}


