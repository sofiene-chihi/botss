package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"e-commerce-chatbot/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConversationTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "conversation.html", gin.H{})
}

func CreateNewConversation(c *gin.Context) {

	systemPrompt := os.Getenv("SYSTEM_PROMPT")

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

	models.AddMessagesToConversation(insertedIDString, conversationMessages)

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
	conversationId := requestData["conversationId"]

	previousContext, err := models.GetConversationById(conversationId)
	if err != nil {
		log.Fatal(err)
	}

	newMessage := models.MessageItem{Role: "user", Content: sentMessage}
	updatedContext := append(previousContext, newMessage)

	botResponse, err := SendBotMessage(updatedContext)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	returnedMessage := string(botResponse.Choices[0].Message.Content)
	newResponse := models.MessageItem{Role: "assistant", Content: returnedMessage}

	models.AddMessagesToConversation(conversationId, []models.MessageItem{newMessage, newResponse})
	c.JSON(http.StatusOK, gin.H{"message": returnedMessage})
}
