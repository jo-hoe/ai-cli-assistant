package config

import (
	"errors"
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"
)

// DefaultPrompt is used when no prompt is provided in configuration.
const DefaultPrompt = `Provide a %s command for the following action: %s.
Provide first an example and afterwards a description. Here is an example how the output scheme should look like:
Command: <example command>

Description: <step by step description of the command>`

type OpenAIConfig struct {
	Endpoint string `yaml:"endpoint"`
	Model    string `yaml:"model"`
}

// Config holds non-secret runtime configuration in YAML.
// Secrets like API keys should continue to be provided via environment variables.
type Config struct {
	// Backend selects the AI provider ("openai", etc.). Defaults to "openai".
	Backend string `yaml:"backend"`
	// CLIKind influences wording (e.g., "bash", "powershell", "kubectl")
	CLIKind string `yaml:"cliKind"`
	// Prompt is the template used to generate the request to the AI backend.
	// Supports %s for CLI kind and action insertion.
	Prompt string `yaml:"prompt"`
	// MaxTokens limits model output tokens
	MaxTokens int `yaml:"maxTokens"`
	// OpenAI specific configuration (used when Backend == "openai")
	OpenAI OpenAIConfig `yaml:"openai"`
}

// Default returns sane defaults.
func Default() Config {
	return Config{
		Backend:  "openai",
		CLIKind:  "cli",
		Prompt:   DefaultPrompt,
		MaxTokens: 256,
		OpenAI: OpenAIConfig{
			Endpoint: "",
			Model:    "",
		},
	}
}

// Load loads configuration from a YAML file.
// If path is empty, it will use the AI_CLI_CONFIG environment variable,
// and if that is also empty, it will default to "config.yaml" in the working directory.
// If the file doesn't exist, Default() is returned without error.
func Load(path string) (Config, error) {
	cfg := Default()

	if path == "" {
		if env := os.Getenv("AI_CLI_CONFIG"); env != "" {
			path = env
		} else {
			path = "config.yaml"
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return cfg, nil
		}
		return cfg, err
	}

	var fileCfg Config
	if err := yaml.Unmarshal(data, &fileCfg); err != nil {
		return cfg, err
	}

	// Overlay defaults with file values (only override when provided)
	if fileCfg.Backend != "" {
		cfg.Backend = fileCfg.Backend
	}
	if fileCfg.CLIKind != "" {
		cfg.CLIKind = fileCfg.CLIKind
	}
	if fileCfg.Prompt != "" {
		cfg.Prompt = fileCfg.Prompt
	}
	if fileCfg.MaxTokens > 0 {
		cfg.MaxTokens = fileCfg.MaxTokens
	}
	if fileCfg.OpenAI.Endpoint != "" {
		cfg.OpenAI.Endpoint = fileCfg.OpenAI.Endpoint
	}
	if fileCfg.OpenAI.Model != "" {
		cfg.OpenAI.Model = fileCfg.OpenAI.Model
	}
	return cfg, nil
}
