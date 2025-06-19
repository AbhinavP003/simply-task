package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"Status": "Running"})
	})
	err := router.Run(":8081")
	if err != nil {
		log.Println("Failed to run user-service")
	}
}
