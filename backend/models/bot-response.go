package models

type BotResponse struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	Index         int         `json:"index"`
	Message       MessageItem `json:"message"`
	Finish_reason string      `json:"finish_reason"`
}

type MessageItem struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	Prompt_tokens     int `json:"prompt_tokens"`
	Completion_tokens int `json:"completion_tokens"`
	Total_tokens      int `json:"total_tokens"`
}
