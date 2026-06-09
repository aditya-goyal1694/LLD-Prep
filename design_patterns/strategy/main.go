package main

import (
	"lld/design_patterns/strategy/checkout"
	"lld/design_patterns/strategy/payment_methods"
)


func main() {
	checkout := checkout.NewCheckout(
		payment_methods.NewUPI("aditya@oksbi", 1234),
	)

	_ = checkout.Pay(1000)

	checkout.SetPaymentMethod(
		payment_methods.NewCreditCard(
			"1234567812345678",
			"Aditya Goyal",
			123,
		),
	)

	_ = checkout.Pay(2500)

	checkout.SetPaymentMethod(
		payment_methods.NewPayPal(
			"aditya@gmail.com",
			"secret-password",
		),
	)

	_ = checkout.Pay(5000)
}