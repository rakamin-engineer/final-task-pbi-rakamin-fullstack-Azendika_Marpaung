package main

import (
	"PROJECT/app"
	"PROJECT/controllers"
	"PROJECT/database"
	"PROJECT/middlewares"

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

		authorized.POST("/photo", controllers.CreatePhoto)
		authorized.GET("/photo", controllers.GetPhoto)
		authorized.PUT("/photo/:photoId", controllers.UpdatePhoto)
		authorized.DELETE("/photo/:photoId", controllers.DeletePhoto)
	}

	r.Run()
}
