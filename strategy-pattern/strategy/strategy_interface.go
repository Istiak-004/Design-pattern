package strategy

// Strategy interface defines a method for executing a strategy.
type PaymentStrategy interface {
	// Pay method executes the payment strategy.
	Pay(amount float64) string
}
