package llm_pricing

import "fmt"

// Model represents a model with its price
type Model struct {
	Provider string `json:"provider"`
	Model    string `json:"model"`
	Currency string `json:"currency"`
	// PriceQuery is the price per token for a query message
	PriceQuery float32 `json:"price_query"`
	// PriceOutput is the price per token for an output message
	PriceOutput float32 `json:"price_output"`
}

// Price is an interface for model price calculation
type Price interface {
	ForModelQuery(provider, model string, currency string, tokens int) (float32, error)
	ForModelOutput(provider, model string, currency string, tokens int) (float32, error)
}

// FormatPrice formats a price with currency
func FormatPrice(price float32, currency string) string {
	switch currency {
	case "USD":
		return fmt.Sprintf("$%.4f", price)
	case "EUR":
		return fmt.Sprintf("€%.4f", price)
	default:
		return fmt.Sprintf("%.4f", price)
	}
}

// Pricing is a list of models with price
type Pricing struct {
	models          []Model
	defaultCurrency string
	currencyRates   map[string]float32
}

var _ Price = (*Pricing)(nil)

// NewPricing returns a new pricing
func NewPricing(models []Model, currencyRates map[string]float32) *Pricing {
	return &Pricing{
		models:        models,
		currencyRates: currencyRates,
	}
}

// ForModelQuery returns the price for a model query
func (p *Pricing) ForModelQuery(provider, model string, currency string, tokens int) (float32, error) {
	for _, m := range p.models {
		if m.Provider == provider && m.Model == model {
			currencyRate, ok := p.currencyRates[currency]
			if !ok {
				return 0, fmt.Errorf("unknown currency %s", currency)
			}

			return m.PriceQuery * float32(tokens) * currencyRate, nil
		}
	}

	return 0.0, fmt.Errorf("unknown model %s", model)
}

// ForModelOutput returns the price for a model output
func (p *Pricing) ForModelOutput(provider, model string, currency string, tokens int) (float32, error) {
	for _, m := range p.models {
		if m.Provider == provider && m.Model == model {
			currencyRate, ok := p.currencyRates[m.Currency]
			if !ok {
				return 0, fmt.Errorf("unknown model currency: %s", currency)
			}

			originalPrice := m.PriceOutput * float32(tokens) * currencyRate

			currencyRate, ok = p.currencyRates[currency]
			if !ok {
				return 0, fmt.Errorf("unknown currency: %s", currency)
			}

			return originalPrice / currencyRate, nil
		}
	}
	return 0.0, fmt.Errorf("unknown model %s", model)
}
