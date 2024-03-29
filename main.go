package main

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hellow world!",
		})
	})

	r.GET("/users", controllers.GetAllUsers)
	r.POST("/users/:id",controllers.GetUser)
	r.Run()
}
