package config

import (
	"context"
	"fmt"
	"log"
	"project/Constants"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDatabase() (*mongo.Client,error){
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	mongoConnection := options.Client().ApplyURI(Constants.ConnectionString)
	mongoClient,err := mongo.Connect(ctx, mongoConnection)
	if err != nil {
		log.Fatal(err.Error())
		return nil,err
	}
	if err := mongoClient.Ping(ctx,readpref.Primary()); err != nil {
		return nil,err
	}
	return mongoClient,nil
}

func Getcollection(client *mongo.Client,dbName string,collectionName string) *mongo.Collection {
	client,err:= ConnectDatabase()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	collection:=client.Database(dbName).Collection(collectionName)
	return collection
}