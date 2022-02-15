package src

import (
	"log"
	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	log.Println("ping requested...")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

