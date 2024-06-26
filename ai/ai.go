package chat_ai

import "fmt"

func Talk(previousPayload *Payload, text string) {
	previousPayload.Messages = append(previousPayload.Messages, PayloadMessage{
		Role:    "user",
		Content: text,
	})

	aiMessage, err := Req(*previousPayload)
	if err != nil {
		panic(err)
	}

	fmt.Println(aiMessage.Text)

	previousPayload.Messages = append(previousPayload.Messages, PayloadMessage{
		Role:    "assistant",
		Content: aiMessage.Text,
	})
}
