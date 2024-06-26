package chat_ai

type PayloadMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Payload struct {
	Model    string           `json:"model"`
	Messages []PayloadMessage `json:"messages"`
}

type ResponseMessage struct {
	Role    string `json:"role"`
	Message string `json:"message"`
}

type Response struct {
	Data []ResponseMessage
	Text string
}
