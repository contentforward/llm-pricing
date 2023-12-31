# LLM Pricing as a package

## Available models

| Provider | Model                      | Price Query | Price Output |
|----------|----------------------------|-------------|--------------|
| openai   | gpt-4-1106-preview         | 0.01        | 0.03         |
| openai   | gpt-4-1106-vision-preview  | 0.01        | 0.03         |
| openai   | gpt-4                      | 0.03        | 0.06         |
| openai   | gpt-4-32k                  | 0.06        | 0.12         |
| openai   | gpt-3.5-turbo-1106         | 0.0010      | 0.0020       |
| openai   | gpt-3.5-turbo-instruct     | 0.0015      | 0.0020       |

## Usage

The Public API is able to calculate the price for a given model and a given number of tokens. 
Token counting is out of scope of this package, but there's [tiktoken-go](https://github.com/pkoukk/tiktoken-go) for that.

It works with multiple currencies and supports conversion between them.

```go
// Price is an interface for model price calculation
type Price interface {
    ForModelQuery(provider, model string, currency string, tokens int) (float32, error)
    ForModelOutput(provider, model string, currency string, tokens int) (float32, error)
}
```

```go
// CurrencyConversion defines the methods that any type of currency converter must implement.
type CurrencyConversion interface {
    Convert(amount CurrencyAmount, fromCurrency, toCurrency string) (CurrencyAmount, error)
}
```

### Golang example

```go

package main

import (
	"fmt"

	llmpricing "github.com/contentforward/llm-pricing"
)

func main() {
	fmt.Printf("%+v", llmpricing.Models) // all models
	fmt.Printf("%+v", llmpricing.OpenAI) // just openai models

	baseCurrency := llmpricing.CurrencyUSD
	rates := map[string]llmpricing.CurrencyAmount{
		baseCurrency:           1,
		llmpricing.CurrencyEUR: 0.91492366,
	}

	conversion := llmpricing.NewConverter(baseCurrency, rates)
	pricing := llmpricing.NewPricing(llmpricing.Models, conversion)
	
	inputTokens := 3200
	inputPrice, err := pricing.ForModelQuery("openai", "gpt-4-1106-preview", llmpricing.CurrencyEUR, inputTokens)
	if err != nil {
		err = fmt.Errorf("failed to get price: %w", err)
		panic(err)
		return
	}

	fmt.Printf("Input price: %s", llmpricing.FormatPrice(inputPrice, llmpricing.CurrencyEUR))

	outputPrice, err := pricing.ForModelOutput("openai", "gpt-4-1106-preview", llmpricing.CurrencyEUR, inputTokens)
	if err != nil {
		err = fmt.Errorf("failed to get price: %w", err)
		panic(err)
		return
	}

	fmt.Printf("Output price: %s", llmpricing.FormatPrice(outputPrice, llmpricing.CurrencyEUR))

```