package implementation

import "strconv"

type CryptoPayment struct {
	// CryptoCurrency is the type of cryptocurrency used for payment.
	CryptoCurrency string
	// WalletAddress is the wallet address for receiving the payment.
	WalletAddress string
}

func (c *CryptoPayment) Pay(amount float64) string {
	// Simulate payment processing with cryptocurrency
	return "Paid " + strconv.FormatFloat(amount, 'f', 4, 64) + " using " + c.CryptoCurrency + " to wallet " + c.WalletAddress
}
