package payment_methods

type PaymentMethod interface {
	Pay(amount int) error
}

type PaymentDetailsValidator interface {
	ValidatePaymentDetails() error
}
