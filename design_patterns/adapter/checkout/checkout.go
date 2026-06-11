package checkout

import "lld/design_patterns/adapter/adapter"

type Checkout struct {
    gateway adapter.PaymentGateway
}

func NewCheckout(gateway adapter.PaymentGateway) *Checkout {
    return &Checkout{
        gateway: gateway,
    }
}

func (c *Checkout) ProcessPayment(amount int) {
    c.gateway.Pay(amount)
}