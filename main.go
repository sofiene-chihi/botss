package main

import (
	"e-commerce-chatbot/handlers"
	"e-commerce-chatbot/models"
	"embed"
	"fmt"
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	//go:embed templates/*
	templatesEmbed embed.FS
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	models.ConnectRedis()
	r := gin.Default()

	stage := os.Getenv("STAGE")
	if stage == "prod" || stage == "staging" {
		fmt.Println("production")
		templ := template.Must(template.New("").ParseFS(
			templatesEmbed, "templates/html/*.html",
		))
		r.SetHTMLTemplate(templ)

	} else {
		r.LoadHTMLGlob("templates/*.html")
		r.Static("/templates/images", "./templates/images")
	}

	r.GET("/", handlers.ConversationTemplate)
	r.POST("/send-message", handlers.SendMessage)

	errors := r.Run(":8080")
	if errors != nil {
		return
	}
}
