package main

import (
	"PROJECT_BTPN/app"
	"PROJECT_BTPN/controllers"
	"PROJECT_BTPN/database"
	"PROJECT_BTPN/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	r := gin.Default()

	r.POST("/users/register", app.RegisterUser)
	r.POST("/users/login", app.LoginUser)

	authorized := r.Group("/")
	authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.PUT("/users/:userId", controllers.UpdateUser)
		authorized.DELETE("/users/:userId", controllers.DeleteUser)

		authorized.POST("/photos", controllers.CreatePhoto)
		authorized.GET("/photos", controllers.GetPhotos)
		authorized.PUT("/photos/:photoId", controllers.UpdatePhoto)
		authorized.DELETE("/photos/:photoId", controllers.DeletePhoto)
	}

	r.Run()
}
