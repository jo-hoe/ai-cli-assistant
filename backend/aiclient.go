package backend

type AIClient interface {
	GetAnswer(prompt string) (string, error)
}
