package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jo-hoe/ai-cli-assistant/internal/app"
	"github.com/jo-hoe/ai-cli-assistant/internal/openai"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Print("Please set the open ai key as environment variable ('OPENAI_API_KEY')")
		return
	}

	result, err := app.Run(os.Args[1:], openai.NewOpenAIClient(apiKey, 256, &http.Client{}))
	if err != nil {
		fmt.Printf("error encountered: %s", err)
	} else {
		fmt.Print(result)
	}
}
