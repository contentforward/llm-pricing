package llm_pricing

// OpenAI is a list of models from OpenAI
var OpenAI = []Model{
	{
		Provider:    "openai",
		Model:       "gpt-4-0125-preview",
		Currency:    CurrencyUSD,
		PriceQuery:  0.01 / 1000,
		PriceOutput: 0.03 / 1000,
	},
	{
		Provider:    "openai",
		Model:       "gpt-4",
		Currency:    CurrencyUSD,
		PriceQuery:  0.03 / 1000,
		PriceOutput: 0.06 / 1000,
	},
	{
		Provider:    "openai",
		Model:       "gpt-4-32k",
		Currency:    CurrencyUSD,
		PriceQuery:  0.06 / 1000,
		PriceOutput: 0.12 / 1000,
	},
	{
		Provider:    "openai",
		Model:       "gpt-3.5-turbo-0125",
		Currency:    CurrencyUSD,
		PriceQuery:  0.0010 / 1000,
		PriceOutput: 0.0020 / 1000,
	},
	{
		Provider:    "openai",
		Model:       "gpt-3.5-turbo-instruct",
		Currency:    CurrencyUSD,
		PriceQuery:  0.0015 / 1000,
		PriceOutput: 0.0020 / 1000,
	},
}

// Models is a list of available models with price
var Models = OpenAI
