package adapter

import (
	"fmt"
	"lld/design_patterns/adapter/fastpay"
)

type PaymentGateway interface {
    Pay(amount int)
}

type PaymentGatewayAdapter struct {
    fastpay *fastpay.FastPay
}

func NewPaymentGatewayAdapter(fastpay *fastpay.FastPay) *PaymentGatewayAdapter {
    return &PaymentGatewayAdapter{
        fastpay: fastpay,
    }
}

func (a *PaymentGatewayAdapter) Pay(amount int) {
    success := a.fastpay.MakePayment(float64(amount))
    if !success {
        fmt.Printf("There was an error in processing payment.")
    }
}