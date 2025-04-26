package ai

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendChatRequestRoute(c *gin.Context) {
	var dto ConversationDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conversation := InitConversation()
	response, err := conversation.askAI(dto.Query)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}
