package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	main.Print()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hellow world!",
		})
	})

	// r.GET("/albums", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"data": Albums,
	// 	})
	// })

	r.Run()
}
