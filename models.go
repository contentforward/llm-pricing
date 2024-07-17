package llm_pricing

// OpenAI is a list of OpenAI models with price
// last updated: 2024-06-23
var OpenAI = []Model{
	{
		Provider:    "openai",
		Model:       "gpt-4o",
		Releases:    []string{"2024-05-13", "*"},
		Currency:    CurrencyUSD,
		PriceQuery:  0.00500 / 1000,
		PriceOutput: 0.01500 / 1000,
	},
	{
		Provider:    "openai",
		Model:       "gpt-4-turbo",
		Releases:    []string{"2024-04-09", "instruct", "*"},
		Currency:    CurrencyUSD,
		PriceQuery:  0.01000 / 1000,
		PriceOutput: 0.03000 / 1000,
	},
	{
		Provider:    "openai",
		Model:       "gpt-4",
		Currency:    CurrencyUSD,
		PriceQuery:  0.03000 / 1000,
		PriceOutput: 0.06000 / 1000,
	},
	{
		Provider:    "openai",
		Model:       "gpt-4-32k",
		Currency:    CurrencyUSD,
		PriceQuery:  0.06000 / 1000,
		PriceOutput: 0.12000 / 1000,
	},
	{
		Provider:    "openai",
		Model:       "gpt-4",
		Releases:    []string{"0125", "0125-preview", "1106-preview", "vision-preview"},
		Currency:    CurrencyUSD,
		PriceQuery:  0.01000 / 1000,
		PriceOutput: 0.03000 / 1000,
	},
	{
		Provider:    "openai",
		Model:       "gpt-3.5-turbo",
		Currency:    CurrencyUSD,
		PriceQuery:  0.00050 / 1000,
		PriceOutput: 0.00150 / 1000,
	},
	{
		Provider:    "openai",
		Model:       "gpt-3.5-turbo",
		Releases:    []string{"0125"},
		Currency:    CurrencyUSD,
		PriceQuery:  0.00050 / 1000,
		PriceOutput: 0.00150 / 1000,
	},
	{
		Provider:    "openai",
		Model:       "gpt-3.5-turbo-instruct",
		Currency:    CurrencyUSD,
		PriceQuery:  0.00150 / 1000,
		PriceOutput: 0.00200 / 1000,
	},
	{
		Provider:    "openai",
		Model:       "gpt-3.5-turbo",
		Releases:    []string{"1106"},
		Currency:    CurrencyUSD,
		PriceQuery:  0.00100 / 1000,
		PriceOutput: 0.00200 / 1000,
	},
	{
		Provider:    "openai",
		Model:       "gpt-3.5-turbo",
		Releases:    []string{"0613"},
		Currency:    CurrencyUSD,
		PriceQuery:  0.00150 / 1000,
		PriceOutput: 0.00200 / 1000,
	},
	{
		Provider:    "openai",
		Model:       "gpt-3.5-turbo",
		Releases:    []string{"16k-0613"},
		Currency:    CurrencyUSD,
		PriceQuery:  0.00300 / 1000,
		PriceOutput: 0.00400 / 1000,
	},
	{
		Provider:    "openai",
		Model:       "gpt-3.5-turbo",
		Releases:    []string{"0301"},
		Currency:    CurrencyUSD,
		PriceQuery:  0.00150 / 1000,
		PriceOutput: 0.00200 / 1000,
	},
}

// Models is a list of available models with price
var Models = OpenAI
