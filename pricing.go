package llm_pricing

// Model represents a model with its price
type Model struct {
	Provider    string  `json:"provider"`
	Model       string  `json:"model"`
	PriceQuery  float32 `json:"price_query"`
	PriceOutput float32 `json:"price_output"`
}

// Price is an interface for model price calculation
type Price interface {
	ForModelQuery(provider, model string, currency string, tokens int) float32
	ForModelOutput(provider, model string, currency string, tokens int) float32
}

// Pricing is a list of models with price
type Pricing struct {
	models          []Model
	defaultCurrency string
	currencyRates   map[string]float32
}

var _ Price = (*Pricing)(nil)

// NewPricing returns a new pricing
func NewPricing(models []Model, defaultCurrency string, currencyRates map[string]float32) *Pricing {
	return &Pricing{
		models:          models,
		defaultCurrency: defaultCurrency,
		currencyRates:   currencyRates,
	}
}

// ForModelQuery returns the price for a model query
func (p *Pricing) ForModelQuery(provider, model string, currency string, tokens int) float32 {
	for _, m := range p.models {
		if m.Provider == provider && m.Model == model {
			f, ok := p.currencyRates[currency]
			if !ok {
				f = p.currencyRates[p.defaultCurrency]
			}
			return m.PriceQuery * float32(tokens) * f
		}
	}

	return 0.0
}

// ForModelOutput returns the price for a model output
func (p *Pricing) ForModelOutput(provider, model string, currency string, tokens int) float32 {
	for _, m := range p.models {
		if m.Provider == provider && m.Model == model {
			f, ok := p.currencyRates[currency]
			if !ok {
				f = p.currencyRates[p.defaultCurrency]
			}
			return m.PriceOutput * float32(tokens) * f
		}
	}
	return 0.0
}
