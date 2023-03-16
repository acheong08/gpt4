package api

type TextCompletion struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
	Error   Error    `json:"error"`
}

type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Param   string `json:"param"`
	Type    string `json:"type"`
}

type Choice struct {
	Text           string          `json:"text"`
	Index          int             `json:"index"`
	Logprobs       interface{}     `json:"logprobs"`
	FinishReason   string          `json:"finish_reason"`
	InternalMetric InternalMetrics `json:"internal_metrics"`
}

type InternalMetrics struct {
	CachedPromptTokens      int `json:"cached_prompt_tokens"`
	TotalAcceptedTokens     int `json:"total_accepted_tokens"`
	TotalBatchedTokens      int `json:"total_batched_tokens"`
	TotalPredictedTokens    int `json:"total_predicted_tokens"`
	TotalRejectedTokens     int `json:"total_rejected_tokens"`
	TotalTokensInCompletion int `json:"total_tokens_in_completion"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
