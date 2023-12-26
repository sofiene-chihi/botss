package handlers

import (
	"bytes"
	"chatbot-store/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func SendBotMessage(context []models.MessageItem) (models.BotResponse, error) {
	API_URL := os.Getenv("OPEN_AI_API")
	// TOKEN := os.Getenv("OPEN_AI_API_KEY")

	REQUEST_URL := fmt.Sprintf("%s/chat/completions", API_URL)
	secretData, err := getApiToken()
	if err != nil {
		log.Fatal("Failed to find token secret in AWS")
	}
	BEARER_TOKEN := fmt.Sprintf("Bearer %s", secretData.OPENAI_API_TOKEN)
	var botResponse models.BotResponse

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

func getApiToken() (models.SecretData, error) {

	secretName := "OpenAI/secrets"
	var secretData models.SecretData
	region := os.Getenv("AWS_REGION")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY_ID"), os.Getenv("SECRET_ACCESS_KEY"), ""),
	},
	)
	if err != nil {
		log.Println("Failed to create session:", err)
		return secretData, err
	}

	svc := secretsmanager.New(sess)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	// Retrieve the secret value
	result, err := svc.GetSecretValue(input)
	if err != nil {
		log.Println("Failed to retrieve secret value:", err)
		return secretData, err
	}

	if result.SecretString != nil {
		// Secret value is a string
		// secretValue := *result.SecretString
		err = json.Unmarshal([]byte(*result.SecretString), &secretData)
		if err != nil {
			fmt.Println("Failed to parse secret value as JSON:", err)
			return secretData, err
		}
		return secretData, nil
	} else {
		log.Println("Secret value not found")
		return secretData, errors.New("Secret value not found")
	}

}
