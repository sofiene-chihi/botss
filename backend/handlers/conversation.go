package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"chatbot-store/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var SystemPrompt string = `
You are OrderBot, an automated service to collect orders for a pizza restaurant. 
You first greet the customer, then collects the order, 
and then asks if it's a pickup or delivery. 
You wait to collect the entire order, then summarize it and check for a final time if the customer wants to add anything else. 
If it's a delivery, you ask for an address. 
Finally you collect the payment.
Make sure to clarify all options, extras and sizes to uniquely identify the item from the menu.
You respond in a short, very conversational friendly style. 
The menu includes 
pepperoni pizza  12.95, 10.00, 7.00 
cheese pizza  10.95, 9.25, 6.50 
eggplant pizza   11.95, 9.75, 6.75 
fries 4.50, 3.50 
greek salad 7.25 
Toppings: 
extra cheese 2.00, 
mushrooms 1.50 
sausage 3.00 
canadian bacon 3.50 
AI sauce 1.50 
peppers 1.00 
Drinks: 
coke 3.00, 2.00, 1.00 
sprite 3.00, 2.00, 1.00 
bottled water 5.00 
`

func GetConversationById(c *gin.Context) {
	
	conversationId := c.Param("id")
	conversationContent, err := models.GetConversationById(conversationId)
	if err != nil {
		fmt.Println("Error retrieving objects from Redis:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not get conversation by ID!!"})
	}
	if len(conversationContent) > 0 {
		// Remove the first element by creating a new slice that excludes it
		filteredConversationContent := conversationContent[1:]
		fmt.Println("Slice after removing the first element:", filteredConversationContent)
		fmt.Println(conversationContent)
		c.JSON(http.StatusOK, gin.H{"conversationContent": filteredConversationContent})
	} else {
		// Remove the first element by creating a new slice that excludes it
		emptyContent := []models.MessageItem{}
		c.JSON(http.StatusOK, gin.H{"conversationContent": emptyContent })
	}
}

func SaveConversation(c *gin.Context) {

	var requestData map[string]string
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(requestData["conversationId"])
	conversationId := requestData["conversationId"]

	models.InsertConversationMessage(conversationId)
	c.JSON(http.StatusOK, gin.H{"Result": "Conversation saved!!"})
}

func CreateNewConversation(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := models.MongoClient.Database("chatbot-conversations").Collection("conversation")

	fmt.Println(SystemPrompt)
	conversationMessages := []models.MessageItem{
		{Role: "system", Content: SystemPrompt},
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
