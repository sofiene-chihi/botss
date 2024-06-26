package main

import (
	"chatbot-store/handlers"
	"chatbot-store/models"
	"embed"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	//go:embed .env
	envFile embed.FS
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	models.ConnectRedis()

	_, err = models.ConnectToMongoDB()
	if err != nil {
		fmt.Println(err)
	}

	r := gin.Default()
	r.Use(corsMiddleware())

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Println(port)

	r.GET("/conversation/:id", handlers.GetConversationById)
	r.POST("/send-message", handlers.SendMessage)
	r.GET("/new-conversation", handlers.CreateNewConversation)
	r.POST("/save-conversation", handlers.SaveConversation)
	r.GET("/healthCheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	errors := r.Run(port)
	if errors != nil {
		return
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
