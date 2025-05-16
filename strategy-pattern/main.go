package main

import (
	"github.com/study/design_pattern/strategy-pattern/context"
	"github.com/study/design_pattern/strategy-pattern/strategy/implementation"
)

func main() {
	// Create a new payment context with a credit card strategy
	creditCardPayment := &implementation.CreditCardPayment{
		CardNumber:     "1234-5678-9012-3456",
		ExpiryDate:     "12/25",
		CVV:            "123",
		CardHolderName: "John Doe",
	}
	paymentContext := context.NewPaymentContext(creditCardPayment)

	// Execute payment
	creditCardPaymentExecution := paymentContext.ExecutePayment(100.0)
	println(creditCardPaymentExecution)

	// Change strategy to PayPal
	paypalPayment := &implementation.PaypalPayment{
		Email:    "test@test.com",
		Password: "password",
		Token:    "test_token",
	}

	paymentContext.SetPaymentStrategy(paypalPayment)
	payPalExecution := paymentContext.ExecutePayment(200.0)
	println(payPalExecution)

}
