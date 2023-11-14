package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"e-commerce-chatbot/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConversationTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "conversation.html", gin.H{})
}

func CreateNewConversation(c *gin.Context) {

	systemPrompt := "hi there, pretend that you're INSAT administrator during our discussion"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := models.MongoClient.Database("chatbot-conversations").Collection("conversation")

	conversationMessages := []models.MessageItem{
		{Role: "system", Content: systemPrompt},
	}
	newConversation := models.Conversation{Messages: conversationMessages}
	insertResult, err := collection.InsertOne(ctx, newConversation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document:", insertResult.InsertedID)

	insertedID, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Fatal("InsertedID is not an ObjectID")
	}

	// Convert ObjectID to string
	insertedIDString := insertedID.Hex()
	fmt.Println("Inserted ID as string:", insertedIDString)

	c.JSON(http.StatusOK, gin.H{"conversationId": insertedIDString})
}

func SendMessage(c *gin.Context) {

	var requestData map[string]string
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(requestData["message"])
	sentMessage := requestData["message"]

	botResponse, err := SendBotMessage(sentMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Print the response body as a string
	returnedMessage := string(botResponse.Choices[0].Message.Content)

	c.JSON(http.StatusOK, gin.H{"message": returnedMessage})
}
