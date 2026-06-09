package payment_methods

import (
	"fmt"
	"strings"
	"errors"
)

type UPI struct {
	id        string
	pin       int
	validator PaymentDetailsValidator
}

func NewUPI(id string, pin int) *UPI {
	upi := &UPI{
		id:  id,
		pin: pin,
	}

	upi.validator = NewUPIValidator(upi)
	return upi
}

func (u *UPI) Pay(amount int) error {
	if err := u.validator.ValidatePaymentDetails(); err != nil {
		return err
	}

	fmt.Printf("Processing payment of ₹%d using UPI ID %s\n", amount, u.id)
	return nil
}

type UPIValidator struct {
	account *UPI
}

func NewUPIValidator(acc *UPI) *UPIValidator {
	return &UPIValidator{
		account: acc,
	}
}

func (u *UPIValidator) ValidatePaymentDetails() error {
	if strings.TrimSpace(u.account.id) == "" {
		return errors.New("upi id cannot be empty")
	}

	if u.account.pin < 1000 || u.account.pin > 9999 {
		return errors.New("invalid upi pin")
	}

	fmt.Printf("UPI account '%s' validated successfully\n",u.account.id)
	return nil
}