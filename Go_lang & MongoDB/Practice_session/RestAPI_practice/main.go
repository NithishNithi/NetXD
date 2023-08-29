package main

import (
	"Rest-API/config"
	"Rest-API/constants"
	"Rest-API/controllers"
	"Rest-API/routes"
	"Rest-API/services"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	ctx  context.Context
	server         *gin.Engine
)
func initRoutes(){
	routes.Default(server)
  }

  func initApp(mongoClient *mongo.Client){
	ctx=context.TODO()
	TransactionCollection:=mongoClient.Database(constants.DatabaseName).Collection("profiles")
	profileService := services.NewTransactionServiceInit(TransactionCollection, ctx)
	profileController := controllers.InitProfileController(profileService)
	routes.ProfileRoute(server,profileController)
  }

func main(){
	server = gin.Default()
	mongoclient,err:=config.ConnectDataBase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	initRoutes()
	initApp(mongoclient)
	fmt.Println("server running on port",constants.Port)
	log.Fatal(server.Run(constants.Port))


}