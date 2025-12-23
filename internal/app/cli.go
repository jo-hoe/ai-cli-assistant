package app

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jo-hoe/ai-cli-assistant/internal/aiclient"
)

const prompt = `Provide a %s command for the following action: %s.
Provide first an example and afterwards a description. Here is an example how the output scheme should look like:
Command: <example command>

Description: <step by step description of the command>`

// Run executes the CLI logic using the provided AI client.
func Run(args []string, aiClient aiclient.AIClient, cliKind string, promptTemplate string) (string, error) {
	if cliKind == "" {
		cliKind = "cli"
	}

	action := strings.Join(args, " ")
	if action == "" {
		return "", errors.New("please define an action to invoke the program")
	}

	tmpl := promptTemplate
	if tmpl == "" {
		tmpl = prompt
	}
	return aiClient.GetAnswer(fmt.Sprintf(tmpl, cliKind, action))
}
