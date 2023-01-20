package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"miniecommerce.wisnu.net/database"
	"miniecommerce.wisnu.net/master/user"
)


func main(){
	r := gin.Default()

	database.ConnectDatabase()
	

	fmt.Println("Connection to database Establish")

	// Initiate Handler
	userHandler := user.User{}

	// Router User
	userRoute := r.Group("/v1/user")
	{
		userRoute.POST("/register",userHandler.Register)
		userRoute.POST("/login",userHandler.Login)
		// userRoute.GET("/get-photo",helpers.Auth(),userHandler.GetPhotoUser)
	}

	// // Router Photo
	// photoHandler := photo.Photo{}
	// photoRoute := r.Group("/v1/photo")
	// {
	// 	photoRoute.POST("/create",helpers.Auth(),photoHandler.CreatePhoto)
	// 	photoRoute.GET("/get-user-photo",helpers.Auth(),photoHandler.GetUserPhoto)
	// }

	// // Router Comment
	// commentHandler := comment.Comment{}
	// commentRoute := r.Group("/v1/comment")
	// {
	// 	commentRoute.POST("/create",helpers.Auth(),commentHandler.InsertComment)
	// 	commentRoute.GET("/get-user-comment",helpers.Auth(),commentHandler.GetCommentByUserId)
	// }
	
	r.Run(":4000") // listen and serve on 0.0.0.0:8080
}