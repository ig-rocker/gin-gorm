package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var Albums = []Album{
	{ID: "001", Title: "Hero Heeralal", Artist: "Amitabh Bachhan", Price: 12.00},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": Albums},
	)
}

func GetUser(c *gin.Context) {
	params := c.Request.URL.Query()
	fmt.Println(params)
	// prints map[]
	// need to bind data

	var data map[string]interface{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}
	fmt.Println(data)
	c.JSON(200, gin.H{
		"message": "data recieved."})
}
