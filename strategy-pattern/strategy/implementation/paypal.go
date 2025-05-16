package implementation

import "strconv"

type PaypalPayment struct {
	// Email is the PayPal account email.
	Email string
	// Password is the PayPal account password.
	Password string
	// Token is the authentication token for PayPal API.
	Token string
}

func (p *PaypalPayment) Pay(amount float64) string {
	// Simulate payment processing with PayPal
	return "Paid " + strconv.FormatFloat(amount, 'f', 4, 64) + " using PayPal account " + p.Email
}
