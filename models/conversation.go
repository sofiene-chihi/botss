package models

type Conversation struct {
	Messages []MessageItem `json:"messages"`
}
