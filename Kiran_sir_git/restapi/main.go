package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	router.GET("/", handleDefault)
	router.Run(":4000")
}

func handleDefault(c *gin.Context){
	c.String(http.StatusOK, "Hello")
}