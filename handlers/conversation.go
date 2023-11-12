package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConversationTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "conversation.html", gin.H{})
}
