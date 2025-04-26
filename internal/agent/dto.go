package agent

type ConversationDto struct {
	Query string `json:"query" binding:"required"`
}

type ConversationResponse struct {
	Response string `json:"response"`
}
