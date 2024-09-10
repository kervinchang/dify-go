package dify

// Usage - Model usage information.
type Usage struct {
	PromptTokens        int     `json:"prompt_tokens"`         // Number of tokens in the prompt.
	PromptUnitPrice     string  `json:"prompt_unit_price"`     // Unit price of the prompt.
	PromptPriceUnit     string  `json:"prompt_price_unit"`     // Price unit of the prompt.
	PromptPrice         string  `json:"prompt_price"`          // Total price of the prompt.
	CompletionTokens    int     `json:"completion_tokens"`     // Number of tokens in the completion.
	CompletionUnitPrice string  `json:"completion_unit_price"` // Unit price of the completion.
	CompletionPriceUnit string  `json:"completion_price_unit"` // Price unit of the completion.
	CompletionPrice     string  `json:"completion_price"`      // Total price of the completion.
	TotalTokens         int     `json:"total_tokens"`          // Total number of tokens used.
	TotalPrice          string  `json:"total_price"`           // Total price of the request.
	Currency            string  `json:"currency"`              // Currency used for the request.
	Latency             float64 `json:"latency"`               // Latency of the request.
}
