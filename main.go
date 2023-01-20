package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"miniecommerce.wisnu.net/database"
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

	// Router User
	userRoute := r.Group("/v1/user")
	{
		userRoute.POST("/register",userHandler.Register)
		userRoute.POST("/login",userHandler.Login)
		// userRoute.GET("/get-photo",helpers.Auth(),userHandler.GetPhotoUser)
	}

	// Router Seller
	sellerRoute := r.Group("/v1/seller")
	{
		sellerRoute.POST("/register",sellerHandler.RegisterSeller)
	}

	// // Router Comment
	// commentHandler := comment.Comment{}
	// commentRoute := r.Group("/v1/comment")
	// {
	// 	commentRoute.POST("/create",helpers.Auth(),commentHandler.InsertComment)
	// 	commentRoute.GET("/get-user-comment",helpers.Auth(),commentHandler.GetCommentByUserId)
	// }
	
	r.Run(":4000") // listen and serve on 0.0.0.0:8080
}