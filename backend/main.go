package main

import (
	"chatbot-store/handlers"
	"chatbot-store/models"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	//go:embed templates/*
	templatesEmbed embed.FS

	//go:embed templates/images
	staticEmbed embed.FS

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


	stage := os.Getenv("STAGE")
	if stage == "prod" {
		fmt.Println("production")
		templ := template.Must(template.New("").ParseFS(
			templatesEmbed, "templates/*.html",
		))
		r.SetHTMLTemplate(templ)
		staticFS, _ := fs.Sub(staticEmbed, "templates/images")
		r.StaticFS("/templates/images", http.FS(staticFS))

	} else {
		r.LoadHTMLGlob("templates/*.html")
		r.Static("/templates/images", "./templates/images")
	}

	r.GET("/", handlers.ConversationTemplate)
	r.POST("/send-message", handlers.SendMessage)
	r.GET("/new-conversation", handlers.CreateNewConversation)
	r.POST("/save-conversation", handlers.SaveConversation)

	errors := r.Run(":8080")
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