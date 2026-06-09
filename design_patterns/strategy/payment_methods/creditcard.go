package payment_methods

import (
	"errors"
	"fmt"
)

type CreditCard struct {
	cardNumber string
	name       string
	cvv        int
	validator  PaymentDetailsValidator
}

func NewCreditCard(cardNumber string, name string, cvv int) *CreditCard {
	card := &CreditCard{
		cardNumber: cardNumber,
		name:       name,
		cvv:        cvv,
	}

	card.validator = NewCreditCardValidator(card)
	return card
}

func (c *CreditCard) Pay(amount int) error {
	if err := c.validator.ValidatePaymentDetails(); err != nil {
		return err
	}

	lastFour := c.cardNumber[len(c.cardNumber)-4:]

	fmt.Printf("Processing payment of ₹%d using Credit Card ending with %s\n", amount, lastFour)
	return nil
}

type CreditCardValidator struct {
	card *CreditCard
}

func NewCreditCardValidator(card *CreditCard) *CreditCardValidator {
	return &CreditCardValidator{
		card: card,
	}
}

func (c *CreditCardValidator) ValidatePaymentDetails() error {
	if len(c.card.cardNumber) != 16 {
		return errors.New("card number must be 16 digits")
	}

	if c.card.cvv < 100 || c.card.cvv > 999 {
		return errors.New("invalid cvv")
	}

	fmt.Printf("Credit Card ending with %s validated successfully\n",c.card.cardNumber[len(c.card.cardNumber)-4:])
	return nil
}