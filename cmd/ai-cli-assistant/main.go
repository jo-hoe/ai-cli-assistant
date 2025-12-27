package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jo-hoe/ai-cli-assistant/internal/app"
	"github.com/jo-hoe/ai-cli-assistant/internal/openai"
	"github.com/jo-hoe/ai-cli-assistant/internal/config"
	"github.com/jo-hoe/ai-cli-assistant/internal/aiclient"
)

func main() {
	// Get default config path in user's home directory
	homeDir, _ := os.UserHomeDir()
	defaultConfigPath := filepath.Join(homeDir, ".ai-cli-assistant", "config.yaml")
	
	// Define command-line flags
	var configPath string
	flag.StringVar(&configPath, "config", defaultConfigPath, "Path to config file")
	flag.StringVar(&configPath, "c", defaultConfigPath, "Path to config file (shorthand)")
	flag.Parse()

	cfg, _ := config.Load(configPath)
	
	// Get API key from environment variable or config file (env var takes precedence)
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		apiKey = cfg.OpenAI.APIKey
	}
	
	if apiKey == "" {
		fmt.Print("Please set the OpenAI API key either:\n- As environment variable 'OPENAI_API_KEY'\n- In config file under 'openai.apiKey'")
		return
	}

	var client aiclient.AIClient
	
	// Check which backend is enabled
	if cfg.OpenAI.Enabled {
		client = openai.NewOpenAIClient(apiKey, cfg.OpenAI.MaxTokens, &http.Client{}, cfg.OpenAI.Endpoint, cfg.OpenAI.Model)
	} else {
		fmt.Print("No AI backend is enabled in configuration")
		return
	}
	result, err := app.Run(flag.Args(), client, cfg.CLIKind, cfg.Prompt)
	if err != nil {
		fmt.Printf("error encountered: %s", err)
	} else {
		fmt.Print(result)
	}
}
