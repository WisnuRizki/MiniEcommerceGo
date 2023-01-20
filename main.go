package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"miniecommerce.wisnu.net/database"
	"miniecommerce.wisnu.net/helpers"
	"miniecommerce.wisnu.net/master/product"
	"miniecommerce.wisnu.net/master/seller"
	"miniecommerce.wisnu.net/master/user"
)


func main(){
	r := gin.Default()

	database.ConnectDatabase()
	

	fmt.Println("Connection to database Establish")

	// Initiate Handler
	userHandler := user.User{}
	sellerHandler := seller.Seller{}
	productHandler := product.Product{}

	// Router User
	userRoute := r.Group("/v1/user")
	{
		userRoute.POST("/register",userHandler.Register)
		userRoute.POST("/login",userHandler.Login)
	}

	// Router Seller
	sellerRoute := r.Group("/v1/seller")
	{
		sellerRoute.POST("/register",sellerHandler.RegisterSeller)
		sellerRoute.POST("/login",sellerHandler.LoginSeller)
		sellerRoute.POST("/create-product",helpers.Auth(),productHandler.Create)
	}

	
	r.Run(":4000") // listen and serve on 0.0.0.0:8080
}