package handlers

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"chatbot-store/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
		c.JSON(http.StatusOK, gin.H{"conversationContent": emptyContent})
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

	systemPrompt, err := getSystemPrompt("pizza-restaurant")

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

func getSystemPrompt(topic string) (string, error) {

	region := os.Getenv("AWS_REGION")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY_ID"), os.Getenv("SECRET_ACCESS_KEY"), ""),
	},
	)
	if err != nil {
		log.Println("Failed to create session:", err)
		return "", err
	}

	bucketName := os.Getenv("PROMPTS_BUCKET")
	downloader := s3manager.NewDownloader(sess)
	promptFileName := ""
	if topic == "pizza-restaurant" {
		promptFileName = "pizza-restaurant-bot.txt"
	} else if topic == "bank" {
		promptFileName = "bank-bot.txt"

	} else if topic == "tech-store" {
		promptFileName = "tech-store-bot.txt"

	} else {
		return "", err
	}

	file, err := os.Create("downloaded-prompt-file.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return "", err
	}
	defer file.Close()

	// Download the file from S3.
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(promptFileName),
		})
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return "", err
	}
	fmt.Println("File downloaded successfully!")

	content, err := ioutil.ReadFile("downloaded-file.txt")
	if err != nil {
		fmt.Println("Error reading file content:", err)
		return "", err
	}

	fmt.Println("File content:")
	fmt.Println(string(content))

	return string(content), nil
}
