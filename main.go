package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/jo-hoe/ai-cli-assistant/backend"
	"github.com/jo-hoe/ai-cli-assistant/backend/openai"
)

const prompt = `Provide a %s command for the following action: %s.
Provide first an example and afterwards a description. Here is an example how the output scheme should look like:
Command: <example command>

Description: <step by step description of the command>`

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Print("Please set the open ai key as environment variable ('OPENAI_API_KEY')")
		return
	}

	result, err := runCliWithAIClient(os.Args[1:], openai.NewOpenAIClient(apiKey, 256, &http.Client{}))

	if err != nil {
		fmt.Printf("error encountered: %s", err)
	} else {
		fmt.Print(result)
	}
}

func runCliWithAIClient(args []string, aiClient backend.AIClient) (string, error) {
	cliKind := os.Getenv("CLI_KIND")
	if cliKind == "" {
		cliKind = "cli"
	}

	action := strings.Join(args, " ")
	if action == "" {
		return "", errors.New("please define an action to invoke the program")
	}

	return aiClient.GetAnswer(fmt.Sprintf(prompt, cliKind, action))
}
