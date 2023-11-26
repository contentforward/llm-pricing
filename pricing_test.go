package llm_pricing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	provider string
	model    string

	currency string
	tokens   int

	expected       float32
	expectedOutput string
}

func Test_Pricing_ForModelOutput(t *testing.T) {
	price := NewPricing(OpenAI, map[string]float32{
		CurrencyUSD: 1.0,
		CurrencyEUR: 0.91,
	})

	tests := []test{
		{
			provider: "openai",
			model:    "gpt-4-1106-preview",

			tokens: 1500,

			currency:       "USD",
			expected:       0.044999998,
			expectedOutput: "$0.0450",
		},
		{
			provider: "openai",
			model:    "gpt-4-1106-preview",

			tokens: 1500,

			currency:       "EUR",
			expected:       0.04095,
			expectedOutput: "€0.0410",
		},
	}

	for _, tt := range tests {
		t.Run(tt.provider+" "+tt.model, func(t *testing.T) {
			actual, err := price.ForModelOutput(tt.provider, tt.model, tt.currency, tt.tokens)
			assert.NoError(t, err)

			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.expectedOutput, FormatPrice(actual, tt.currency))
		})
	}
}
