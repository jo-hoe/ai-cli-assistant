# AI CLI Assistant

[![Go Reference](https://pkg.go.dev/badge/jo-hoe/ai-cli-assistant.svg)](https://pkg.go.dev/github.com/jo-hoe/ai-cli-assistant)
[![Test Status](https://github.com/jo-hoe/ai-cli-assistant/workflows/test/badge.svg)](https://github.com/jo-hoe/ai-cli-assistant/actions?workflow=test)
[![Coverage Status](https://coveralls.io/repos/github/jo-hoe/ai-cli-assistant/badge.svg?branch=main)](https://coveralls.io/github/jo-hoe/ai-cli-assistant?branch=main)
[![Lint Status](https://github.com/jo-hoe/ai-cli-assistant/workflows/lint/badge.svg)](https://github.com/jo-hoe/ai-cli-assistant/actions?workflow=lint)
[![CodeQL Status](https://github.com/jo-hoe/ai-cli-assistant/workflows/CodeQL/badge.svg)](https://github.com/jo-hoe/ai-cli-assistant/actions?workflow=CodeQL)
[![Go Report Card](https://goreportcard.com/badge/github.com/jo-hoe/ai-cli-assistant)](https://goreportcard.com/report/github.com/jo-hoe/ai-cli-assistant)

Provides CLI commands based on natural language using GenAI.
The tool is platform and CLI independent (works on Mac, Windows, and Linux).

![Demo](resources/demo_powershell.gif)

## Prerequisites

- Go 1.20+
- An [OpenAI API key](https://platform.openai.com/account/api-keys)


## Configuration (YAML)

Non-secret configuration is read from a YAML config file. Secrets (API keys) are provided via environment variables.

- Config lookup order:
  1) If the environment variable AI_CLI_CONFIG is set, that path is used
  2) Otherwise, config.yaml in the current working directory is used
  3) If no file is found, built-in defaults are used

Supported schema:
```yaml
# Select the AI backend provider. Currently supported: "openai"
backend: openai

# Influences wording (e.g., "bash", "powershell", "kubectl")
cliKind: cli

# Optional prompt template. If omitted, a sensible default is used.
# The template should contain two %s placeholders: the first for cliKind, the second for the action text.
# prompt: |
#   Provide a %s command for the following action: %s.
#   Provide first an example and afterwards a description. Here is an example how the output scheme should look like:
#   Command: <example command>
#
#   Description: <step by step description of the command>

# Maximum number of tokens for the response
maxTokens: 256

# OpenAI backend-specific configuration
openai:
  # Optional: override the default endpoint (e.g., to route via a local proxy)
  # Leave empty to use the default: https://api.openai.com/v1/chat/completions
  endpoint: ""
  # Optional: override the default model (e.g., "gpt-4o-mini", "gpt-3.5-turbo")
  # Leave empty to use the default model.
  model: ""
```

You can find a template at config.example.yaml. Copy it to config.yaml and adjust as needed, or set AI_CLI_CONFIG to point to a specific file:

- Windows PowerShell: `$Env:AI_CLI_CONFIG = "C:\path\to\my-config.yaml"`
- bash/zsh: `export AI_CLI_CONFIG="/path/to/my-config.yaml"`

Secrets:

- OPENAI_API_KEY must be provided via environment variable (not in the config file)

## Build

Build the CLI binary from the entrypoint package:

```bash
go build -o ai-cli-assistant ./cmd/ai-cli-assistant
```

## Run

Set your OpenAI API key as environment variable OPENAI_API_KEY.

Windows (PowerShell):

```powershell
$Env:OPENAI_API_KEY = "sk-..."
# optional: set CLI kind to tailor wording (e.g., powershell, bash, kubectl)
$Env:CLI_KIND = "powershell"

.\ai-cli-assistant "list all folders in this folder"
```

macOS/Linux (bash/zsh):

```bash
export OPENAI_API_KEY="sk-..."
# optional:
export CLI_KIND="bash"

./ai-cli-assistant "list all folders in this folder"
```

You can also run without building:

```bash
go run ./cmd/ai-cli-assistant "list all folders in this folder"
```

## Testing

Run all tests:

```bash
go test ./...
```
