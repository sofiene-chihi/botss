package handlers

import (
	"bytes"
	"e-commerce-chatbot/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func SendBotMessage(context []models.MessageItem) (models.BotResponse, error) {
	API_URL := os.Getenv("OPEN_AI_API")
	TOKEN := os.Getenv("OPEN_AI_API_KEY")

	REQUEST_URL := fmt.Sprintf("%s/chat/completions", API_URL)
	BEARER_TOKEN := fmt.Sprintf("Bearer %s", TOKEN)
	var botResponse models.BotResponse

	// systemPrompt := os.Getenv("SYSTEM_PROMPT")

	payload := map[string]interface{}{
		"model":       "gpt-3.5-turbo",
		"messages":    context,
		"temperature": 0.7,
	}

	// Convert payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return botResponse, err
	}

	fmt.Println(REQUEST_URL)
	request, err := http.NewRequest("POST", REQUEST_URL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return botResponse, err
	}

	// Set headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", BEARER_TOKEN)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making request:", err)
		return botResponse, err
	}
	defer response.Body.Close()
	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return botResponse, err
	}
	// fmt.Println(string(responseBody))

	error := json.Unmarshal(responseBody, &botResponse)
	if error != nil {
		fmt.Println("Error decoding JSON:", error)
		return botResponse, err
	}

	return botResponse, nil
}
