package main

import (
    "lld/design_patterns/adapter/adapter"
    "lld/design_patterns/adapter/checkout"
    "lld/design_patterns/adapter/fastpay"
)

func main() {
    fastpay := &fastpay.FastPay{}
	adapter := adapter.NewPaymentGatewayAdapter(fastpay)
    checkout := checkout.NewCheckout(adapter)
    checkout.ProcessPayment(100)
}