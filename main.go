package main

import (
	"e-commerce-chatbot/handlers"
	"embed"
	"fmt"
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	//go:embed templates/*
	templatesEmbed embed.FS

	env = os.Getenv("GO_ENV")
)

func main() {
	fmt.Println("Hello to e-commerce chatbot")
	r := gin.Default()

	if env == "prod" || env == "dev" {
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
	errors := r.Run(":8080")
	if errors != nil {
		return
	}
}
