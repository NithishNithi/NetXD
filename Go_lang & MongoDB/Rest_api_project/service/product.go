package service

import (
	"context"
	"fmt"
	"project/config"
	"project/models"
	"time"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// func InsertProduct(){
// 	var product models.Product
// 	product.ID=primitive.NewObjectID()
// 	product.Name="Iphone"
// 	product.Price = 115000
// 	product.Description = "its an iphone"

// 	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
// 	defer cancel()
// 	client,_:=config.ConnectDatabase()
// 	var productCollection *mongo.Collection = config.Getcollection(client,"inventory","products")
// 	result,err := productCollection.InsertOne(ctx,product)
// 	if err != nil{
// 		fmt.Println(err)
// 	}
// 	fmt.Println(result)
// }
func ProductContext()*mongo.Collection {
	client,_ := config.ConnectDatabase()
	return config.Getcollection(client,"inventory","products")
}

func InsertProduct(product models.Product)(*mongo.InsertOneResult,error) {
	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
	result,err:=ProductContext().InsertOne(ctx,product)
	if err!= nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return result,nil
}
func InsertProductList(products []interface{})(*mongo.InsertManyResult,error) {
	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
	result,err:=ProductContext().InsertMany(ctx,products)
	if err!= nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return result,nil
}

func FindProducts()([]*models.Product,error){
    ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
    filter:=bson.D{{"name","samsung"}}
    result,err:=ProductContext().Find(ctx,filter)
    if err!=nil{
        fmt.Println(err.Error())
        return nil,err
    }else{
        fmt.Println(result)
        var products [] *models.Product
        for result.Next(ctx){
            product:=&models.Product{}
            err:=result.Decode(product)
            if err!=nil{
                return nil,err
            }
            fmt.Println(products)
            products=append(products, product)
        }
        if err:=result.Err(); err!=nil{
            return nil,err
        }
        if len(products)==0{
              return []*models.Product{}, nil
        }
        return products,nil
    }
}