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
	Enabled   bool   `yaml:"enabled"`
	Endpoint  string `yaml:"endpoint"`
	Model     string `yaml:"model"`
	MaxTokens int    `yaml:"maxTokens"`
}

// Config holds non-secret runtime configuration in YAML.
// Secrets like API keys should continue to be provided via environment variables.
type Config struct {
	// CLIKind influences wording (e.g., "bash", "powershell", "kubectl")
	CLIKind string `yaml:"cliKind"`
	// Prompt is the template used to generate the request to the AI backend.
	// Supports %s for CLI kind and action insertion.
	Prompt string `yaml:"prompt"`
	// OpenAI specific configuration
	OpenAI OpenAIConfig `yaml:"openai"`
}

// Default returns sane defaults.
func Default() Config {
	return Config{
		CLIKind: "cli",
		Prompt:  DefaultPrompt,
		OpenAI: OpenAIConfig{
			Enabled:   true,
			Endpoint:  "",
			Model:     "",
			MaxTokens: 256,
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
	if fileCfg.CLIKind != "" {
		cfg.CLIKind = fileCfg.CLIKind
	}
	if fileCfg.Prompt != "" {
		cfg.Prompt = fileCfg.Prompt
	}
	// OpenAI config: merge settings
	cfg.OpenAI.Enabled = fileCfg.OpenAI.Enabled
	if fileCfg.OpenAI.Endpoint != "" {
		cfg.OpenAI.Endpoint = fileCfg.OpenAI.Endpoint
	}
	if fileCfg.OpenAI.Model != "" {
		cfg.OpenAI.Model = fileCfg.OpenAI.Model
	}
	if fileCfg.OpenAI.MaxTokens > 0 {
		cfg.OpenAI.MaxTokens = fileCfg.OpenAI.MaxTokens
	}
	return cfg, nil
}
