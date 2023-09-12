package main

import (
	tokenvalidation "Restapi_k6/token_validation"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Item struct {
	Cusid     int    `json:"cusid"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Bankid    string `json:"bankid"`
	Balance   int    `json:"balance"`
	Isactive  bool   `json:"isactive"`
}

func homePage(c *gin.Context) {

	token := c.GetHeader("Authorization")
	_, err1 := tokenvalidation.ValidateToken(token)
	if err1 != nil {
		return
	}

	var result Item

	if err := c.ShouldBindJSON(&result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := collection.InsertOne(context.Background(), result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data inserted successfully"})
}

func handleRequests() {
	r := gin.Default()
	r.POST("/api", homePage)
	r.Run(":8080")
}

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://mrnithisht:mrnithisht@banking.taaq05t.mongodb.net/") // Update with your MongoDB URI.
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("test").Collection("items")
}

func main() {
	fmt.Println("Server started on port 8080")
	handleRequests()
}
