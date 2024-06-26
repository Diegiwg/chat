package main

import (
	ai "github.com/Diegiwg/chat/ai"
)

func main() {
	history := &ai.Payload{
		Model:    "mistralai/Mixtral-8x7B-Instruct-v0.1",
		Messages: []ai.PayloadMessage{},
	}

	ai.Talk(history, "Hello, how are you?")
}
