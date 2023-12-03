package factory

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64) string
}

type CashPM struct{}

func (c *CashPM) Pay(amount float64) string {
	return fmt.Sprintf("Cash payment method is used to pay %f", amount)
}

type CreditCardPM struct{}

func (c *CreditCardPM) Pay(amount float64) string {
	return fmt.Sprintf("Credit Card payment method is used to pay %f", amount)
}

const (
	Cash       = 1
	CreditCard = 2
)

func GetPaymentMethod(paymentMethod int) (PaymentMethod, error) {
	switch paymentMethod {
	case Cash:
		return new(CashPM), nil
	case CreditCard:
		return new(CreditCardPM), nil
	default:
		return nil, errors.New("input payment method does not exist")
	}
}
