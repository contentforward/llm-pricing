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
	case CurrencyUSD:
		return fmt.Sprintf("$%.4f", price)
	case CurrencyEUR:
		return fmt.Sprintf("€%.4f", price)
	default:
		return fmt.Sprintf("%.4f", price)
	}
}

// Pricing is a list of models with price
type Pricing struct {
	models    []Model
	converter Converter
}

var _ Price = (*Pricing)(nil)

// NewPricing returns a new pricing
func NewPricing(models []Model, converter Converter) *Pricing {
	return &Pricing{
		models:    models,
		converter: converter,
	}
}

// ForModelQuery returns the price for a model query
func (p *Pricing) ForModelQuery(provider, model string, currency string, tokens int) (float32, error) {
	return p.forModel(provider, model, currency, tokens, func(m Model) float32 {
		return m.PriceQuery
	})
}

// ForModelOutput returns the price for a model output
func (p *Pricing) ForModelOutput(provider, model string, currency string, tokens int) (float32, error) {
	return p.forModel(provider, model, currency, tokens, func(m Model) float32 {
		return m.PriceOutput
	})
}

func (p *Pricing) forModel(provider, model string, userCurrency string, tokens int, pricePerTokenValueFunc func(Model) float32) (float32, error) {
	for _, m := range p.models {
		if m.Provider != provider || m.Model != model {
			continue
		}

		price := float32(tokens) * pricePerTokenValueFunc(m)
		convert, err := p.converter.Convert(CurrencyAmount(price), m.Currency, userCurrency)
		if err != nil {
			return 0.0, fmt.Errorf("failed to convert price from %s to %s: %w", m.Currency, userCurrency, err)
		}

		return float32(convert), nil
	}

	return 0.0, fmt.Errorf("unknown model %s", model)
}
