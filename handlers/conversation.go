package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConversationTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "conversation.html", gin.H{})
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
