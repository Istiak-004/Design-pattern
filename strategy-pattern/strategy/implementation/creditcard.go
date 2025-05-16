package implementation

import "strconv"

type CreditCardPayment struct {
	// CardNumber is the credit card number.
	CardNumber string
	// ExpiryDate is the expiration date of the credit card.
	ExpiryDate string
	// CVV is the card verification value.
	CVV string
	// CardHolderName is the name of the cardholder.
	CardHolderName string
}

func (c *CreditCardPayment) Pay(amount float64) string {
	// Simulate payment processing with credit card
	return "Paid " + strconv.FormatFloat(amount, 'f', 4, 64) + " using credit card " + c.CardNumber + " by " + c.CardHolderName
}
