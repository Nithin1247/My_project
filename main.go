package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	ctx := context.Background()

	llm, err := ollama.New(ollama.WithModel("llama3.2:3b"))
	if err != nil {
		log.Fatalf("Failed to initialize Ollama: %v", err)
	}

	// Your code gets the real time, so the model doesn't have to
	now := time.Now().Format("Monday, Jan 2, 2006 at 3:04PM")

	userQuestion := "Hey AI! What day of the week is it today, and what time is it right now?"

	prompt := fmt.Sprintf(
		"The current date and time is: %s.\nUse this information to answer the user's question.\n\nQuestion: %s",
		now, userQuestion,
	)

	fmt.Printf("User Prompt: %s\n\nThinking...\n", userQuestion)

	answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatalf("Model error: %v", err)
	}

	fmt.Println("\n--- Answer ---")
	fmt.Println(answer)
}