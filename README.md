# AI CLI Assistant

[![Go Reference](https://pkg.go.dev/badge/jo-hoe/ai-cli-assistant.svg)](https://pkg.go.dev/github.com/jo-hoe/ai-cli-assistant)
[![Test Status](https://github.com/jo-hoe/ai-cli-assistant/workflows/test/badge.svg)](https://github.com/jo-hoe/ai-cli-assistant/actions?workflow=test)
[![Coverage Status](https://coveralls.io/repos/github/jo-hoe/ai-cli-assistant/badge.svg?branch=main)](https://coveralls.io/github/jo-hoe/ai-cli-assistant?branch=main)
[![Lint Status](https://github.com/jo-hoe/ai-cli-assistant/workflows/lint/badge.svg)](https://github.com/jo-hoe/ai-cli-assistant/actions?workflow=lint)
[![CodeQL Status](https://github.com/jo-hoe/ai-cli-assistant/workflows/CodeQL/badge.svg)](https://github.com/jo-hoe/ai-cli-assistant/actions?workflow=CodeQL)
[![Go Report Card](https://goreportcard.com/badge/github.com/jo-hoe/ai-cli-assistant)](https://goreportcard.com/report/github.com/jo-hoe/ai-cli-assistant)

Provides CLI commands based on natural language using GenAI.
The tool is platform and cli independent (works on Mac, Windows, and Linux).

![Demo](resources/demo_powershell.gif)

## Prerequisites

- have an [OpenAI API](https://platform.openai.com/account/api-keys) key
- install golang
- build the binary by navigating into this folder and executing

```terminal
go build .
```

or by defining your customer binary name via

```powershell
go build -o <your desired program name>
```

## Execution

Set your OpenAI API key as environment variable `OPENAI_API_KEY`.
You may set it in powershell via this command:

```powershell
$Env:OPENAI_API_KEY = "sk-w0Ak...."
```

You can either define the available cli tools or environment in the prompt or set optionally it via the  `CLI_KIND`.

```powershell
$Env:CLI_KIND = "powershell"
```

Execute the program by running

```powershell
.\<your program name> list all folder in this folder`
```

## Cost

I currently pay ~0.01 cent for 20 commands.
Although your results may vary, there is a max [token](https://platform.openai.com/tokenizer) limit (256) set in the tool to ensure you do not overspend.
In addition you can set monthly max spending limits in the [OpenAI website](https://platform.openai.com/account/billing/limits).
