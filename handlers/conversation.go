package handlers

import (
	"bytes"
	"e-commerce-chatbot/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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

	API_URL := fmt.Sprintf("%s%s", os.Getenv("OPEN_AI_API"), "/chat/completions")
	BEARER_TOKEN := fmt.Sprintf("%s%s", "Bearer ", os.Getenv("OPEN_AI_API_KEY"))

	payload := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "user", "content": sentMessage},
		},
		"temperature": 0.7,
	}

	// Convert payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	request, err := http.NewRequest("POST", API_URL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", BEARER_TOKEN)

	client := &http.Client{}

	// Make the request
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()
	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println(string(responseBody))

	var botResponse models.BotResponse

	error := json.Unmarshal(responseBody, &botResponse)
	if error != nil {
		fmt.Println("Error decoding JSON:", error)
		return
	}
	// Print the response body as a string
	returnedResponse := string(botResponse.Choices[0].Message.Content)

	c.JSON(http.StatusOK, gin.H{"message": returnedResponse})
}
