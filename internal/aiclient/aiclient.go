package aiclient

// AIClient is the interface for any AI backend capable of returning an answer for a given prompt.
type AIClient interface {
	GetAnswer(prompt string) (string, error)
}
