package checkout

import "lld/design_patterns/strategy/payment_methods"

type Checkout struct {
	paymentMethod payment_methods.PaymentMethod
}

func NewCheckout(pm payment_methods.PaymentMethod) *Checkout {
	return &Checkout{
		paymentMethod: pm,
	}
}

func (c *Checkout) SetPaymentMethod(pm payment_methods.PaymentMethod) {
	c.paymentMethod = pm
}

func (c *Checkout) Pay(amount int) error {
	return c.paymentMethod.Pay(amount)
}