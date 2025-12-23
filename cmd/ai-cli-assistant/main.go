package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jo-hoe/ai-cli-assistant/internal/app"
	"github.com/jo-hoe/ai-cli-assistant/internal/openai"
	"github.com/jo-hoe/ai-cli-assistant/internal/config"
	"github.com/jo-hoe/ai-cli-assistant/internal/aiclient"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Print("Please set the open ai key as environment variable ('OPENAI_API_KEY')")
		return
	}

	cfg, _ := config.Load("")
	var client aiclient.AIClient
	switch cfg.Backend {
	case "openai":
		client = openai.NewOpenAIClient(apiKey, cfg.MaxTokens, &http.Client{}, cfg.OpenAI.Endpoint, cfg.OpenAI.Model)
	default:
		fmt.Printf("unsupported backend: %s", cfg.Backend)
		return
	}
	result, err := app.Run(os.Args[1:], client, cfg.CLIKind, cfg.Prompt)
	if err != nil {
		fmt.Printf("error encountered: %s", err)
	} else {
		fmt.Print(result)
	}
}
