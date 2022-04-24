package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	err := r.Run(":3000")
	if err != nil {
		log.Panic(fmt.Sprintf("Error initializing server: %s", err))
	}
}
