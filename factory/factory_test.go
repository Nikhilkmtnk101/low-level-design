package factory

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	paymentMethod, err := GetPaymentMethod(Cash)
	assert.Nil(t, err, "Cash payment method must exist")
	_, ok := paymentMethod.(*CashPM)
	assert.Equal(t, true, ok,
		"GetPaymentMethod returns different payment method while calling for cash")
}

func TestGetPaymentMethodCreditCard(t *testing.T) {
	paymentMethod, err := GetPaymentMethod(CreditCard)
	assert.Nil(t, err, "Credit Card payment method must exist")
	_, ok := paymentMethod.(*CreditCardPM)
	assert.Equal(t, true, ok,
		"GetPaymentMethod returns different payment method while calling for credit Card")
}

func TestCashPM_Pay(t *testing.T) {
	paymentMethod, err := GetPaymentMethod(Cash)
	assert.Nil(t, err, "Cash payment method must exist")
	msg := paymentMethod.Pay(100)
	if !strings.Contains(msg, "Cash payment method is used") {
		t.Errorf("Pay method in cash payment method has not implemented properly")
	}
}

func TestCreditCardPM_Pay(t *testing.T) {
	paymentMethod, err := GetPaymentMethod(CreditCard)
	assert.Nil(t, err, "Credit Card payment method must exist")
	msg := paymentMethod.Pay(100)
	if !strings.Contains(msg, "Credit Card payment method is used") {
		t.Errorf("Pay method in credit card payment method has not implemented properly")
	}
}
