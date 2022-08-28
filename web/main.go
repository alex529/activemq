package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := MakeConfig("./config.yaml")
	if err != nil {
		log.Panic("could not load configuration")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/bye", func(c *gin.Context) {
		log.Println("Fuck you")
		c.JSON(http.StatusOK, gin.H{
			"message": "bye, bye",
		})
	})

	r.POST("/hello", func(c *gin.Context) {
		var msg struct {
			Name, Notification string
		}
		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		log.Printf("Hello message received form %s, and says %s\n", msg.Name, msg.Notification)

		c.JSON(http.StatusOK, gin.H{"status": "received"})
	})

	r.POST("/hello2", func(c *gin.Context) {
		var msg struct {
			Name, Notification string
		}
		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		log.Printf("Hello2 message received form %s, and says %s\n", msg.Name, msg.Notification)

		c.JSON(http.StatusOK, gin.H{"status": "received"})
	})

	r.Run(":" + cfg.Server.Port)
}
