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
	
	var client aiclient.AIClient
	
	// Check which backend is enabled
	if cfg.OpenAI.Enabled {
		openaiClient, err := openai.NewOpenAIClient(cfg.OpenAI.APIKey, cfg.OpenAI.MaxTokens, &http.Client{}, cfg.OpenAI.Endpoint, cfg.OpenAI.Model)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		client = openaiClient
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
