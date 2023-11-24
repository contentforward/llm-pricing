package llm_pricing

// Model represents a model with its price
type Model struct {
	Provider    string  `json:"provider"`
	Model       string  `json:"model"`
	PriceQuery  float64 `json:"price_query"`
	PriceOutput float64 `json:"price_output"`
}

// OpenAIModels is a list of models from OpenAI
var OpenAIModels = []Model{
	{
		Provider:    "openai",
		Model:       "gpt-4-1106-preview",
		PriceQuery:  0.01,
		PriceOutput: 0.03,
	},
	{
		Provider:    "openai",
		Model:       "gpt-4-1106-vision-preview",
		PriceQuery:  0.01,
		PriceOutput: 0.03,
	},
	{
		Provider:    "openai",
		Model:       "gpt-4",
		PriceQuery:  0.03,
		PriceOutput: 0.06,
	},
	{
		Provider:    "openai",
		Model:       "gpt-4-32k",
		PriceQuery:  0.06,
		PriceOutput: 0.12,
	},
	{
		Provider:    "openai",
		Model:       "gpt-3.5-turbo-1106",
		PriceQuery:  0.0010,
		PriceOutput: 0.0020,
	},
	{
		Provider:    "openai",
		Model:       "gpt-3.5-turbo-instruct",
		PriceQuery:  0.0015,
		PriceOutput: 0.0020,
	},
}

// Models is a list of available models with price
var Models = OpenAIModels
