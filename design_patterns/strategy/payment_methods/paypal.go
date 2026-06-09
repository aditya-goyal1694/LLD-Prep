package payment_methods

import (
	"errors"
	"fmt"
	"strings"
)

type PayPal struct {
	username  string
	password  string
	validator PaymentDetailsValidator
}

func NewPayPal(username string, password string) *PayPal {
	p := &PayPal{
		username: username,
		password: password,
	}

	p.validator = NewPayPalValidator(p)
	return p
}

func (p *PayPal) Pay(amount int) error {
	if err := p.validator.ValidatePaymentDetails(); err != nil {
		return err
	}

	fmt.Printf("Processing payment of ₹%d using PayPal account %s\n", amount, p.username)
	return nil
}

type PayPalValidator struct {
	account *PayPal
}

func NewPayPalValidator(acc *PayPal) *PayPalValidator {
	return &PayPalValidator{
		account: acc,
	}
}

func (p *PayPalValidator) ValidatePaymentDetails() error {
	if strings.TrimSpace(p.account.username) == "" {
		return errors.New("username cannot be empty")
	}

	if strings.TrimSpace(p.account.password) == "" {
		return errors.New("password cannot be empty")
	}

	fmt.Printf("PayPal account '%s' validated successfully\n",p.account.username)

	return nil
}
