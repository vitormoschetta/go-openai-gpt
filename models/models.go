package models

type CompletionRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
	N           int     `json:"n"`
}

type AppSettings struct {
	Token             string            `json:"token"`
	Url               string            `json:"url"`
	CompletionRequest CompletionRequest `json:"completion"`
}

type CompletionResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string    `json:"text"`
		Index        int       `json:"index"`
		Logprobs     *struct{} `json:"logprobs"`
		FinishReason string    `json:"finish_reason"`
	} `json:"choices"`
}
