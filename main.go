package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"miniecommerce.wisnu.net/database"
	"miniecommerce.wisnu.net/helpers"
	"miniecommerce.wisnu.net/master/admin"
	"miniecommerce.wisnu.net/master/balance"
	"miniecommerce.wisnu.net/master/category"
	"miniecommerce.wisnu.net/master/payment"
	"miniecommerce.wisnu.net/master/product"
	"miniecommerce.wisnu.net/master/seller"
	"miniecommerce.wisnu.net/master/user"
)

// Todo
// Buat Table Order
// Buat Tabel Transaction
// Buat End Point Callback



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
	categoryHandler := category.Category{}

	// Router User
	userRoute := r.Group("/v1/user")
	{
		userRoute.POST("/register",userHandler.Register)
		userRoute.POST("/login",userHandler.Login)
		userRoute.POST("/buy-product",helpers.Auth(),userHandler.BuyProduct)
		userRoute.GET("/history",helpers.Auth(),userHandler.CheckHistory)
		userRoute.GET("/balance",helpers.Auth(),userHandler.CheckBalance)
	}

	// Router Seller
	sellerRoute := r.Group("/v1/seller")
	{
		sellerRoute.POST("/register",sellerHandler.RegisterSeller)
		sellerRoute.POST("/login",sellerHandler.LoginSeller)
		sellerRoute.POST("/create-product",helpers.Auth(),productHandler.Create)
		sellerRoute.GET("/get-product",helpers.Auth(),productHandler.GetAllProductBySeller)
		sellerRoute.DELETE("/delete-product/:id",helpers.Auth(),productHandler.DeleteProductSeller)
	}

	// Router Admin
	adminRoute := r.Group("/v1/admin")
	{
		adminRoute.POST("/register",adminHandler.RegisterAdmin)
		adminRoute.POST("/login",adminHandler.LoginAdmin)
		adminRoute.PUT("/add-balance",helpers.Auth(),balanceHandler.AddBalanceUser)
		adminRoute.POST("/add-category",helpers.Auth(),categoryHandler.CreateCategory)
		adminRoute.DELETE("/delete-category/:id",helpers.Auth(),categoryHandler.DeleteCategory)
		adminRoute.GET("/get-category/",helpers.Auth(),categoryHandler.GetAllCategory)
		adminRoute.GET("/get-category-id/:id",helpers.Auth(),categoryHandler.GetCategoryById)
		adminRoute.PUT("/update-category/:id",helpers.Auth(),categoryHandler.UpdateCategoryById)
		adminRoute.DELETE("/delete-user/:id",helpers.Auth(),userHandler.DeleteUser)
		adminRoute.GET("/get-all-balance/",helpers.Auth(),balanceHandler.GetAllBalance)
	}

	callBackRoute := r.Group("/v1/callback")
	{
		callBackRoute.POST("/callback-midtrans",payment.CallBackMidtrans)
	}

	
	r.Run(":4000") // listen and serve on 0.0.0.0:8080
}