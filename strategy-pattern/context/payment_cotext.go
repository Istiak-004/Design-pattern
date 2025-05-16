package context

import "github.com/study/design_pattern/strategy-pattern/strategy"

type PaymentContext struct {
	// PaymentStrategy is the strategy used for payment processing.
	PaymentStrategy strategy.PaymentStrategy
}

// NewPaymentContext creates a new PaymentContext with the specified payment strategy.
func NewPaymentContext(strategy strategy.PaymentStrategy) *PaymentContext {
	return &PaymentContext{
		PaymentStrategy: strategy,
	}
}

// ExecutePayment executes the payment using the specified strategy.
func (p *PaymentContext) ExecutePayment(amount float64) string {
	return p.PaymentStrategy.Pay(amount)
}

// SetPaymentStrategy sets a new payment strategy for the context.
func (p *PaymentContext) SetPaymentStrategy(strategy strategy.PaymentStrategy) {
	p.PaymentStrategy = strategy
}
