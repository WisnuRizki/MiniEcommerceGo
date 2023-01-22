package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"miniecommerce.wisnu.net/database"
	"miniecommerce.wisnu.net/helpers"
	"miniecommerce.wisnu.net/master/admin"
	"miniecommerce.wisnu.net/master/balance"
	"miniecommerce.wisnu.net/master/product"
	"miniecommerce.wisnu.net/master/seller"
	"miniecommerce.wisnu.net/master/user"
)

// Todo
// [V] Register Admin
// [V] Login Admin
// [] Add Balance
// [] Add Category



func main(){
	r := gin.Default()

	database.ConnectDatabase()
	

	fmt.Println("Connection to database Establish")

	// Initiate Handler
	userHandler := user.User{}
	sellerHandler := seller.Seller{}
	productHandler := product.Product{}
	adminHandler := admin.Admin{}
	balanceHandler := balance.Balance{}

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

	// Router Admin
	adminRoute := r.Group("/v1/admin")
	{
		adminRoute.POST("/register",adminHandler.RegisterAdmin)
		adminRoute.POST("/login",adminHandler.LoginAdmin)
		adminRoute.PUT("/add-balance",helpers.Auth(),balanceHandler.AddBalanceUser)
	}

	
	r.Run(":4000") // listen and serve on 0.0.0.0:8080
}