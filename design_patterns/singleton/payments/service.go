package payments

import (
	"lld/design_patterns/singleton/logger"
)

type PaymentService struct{}

func (p *PaymentService) PrintPaymentDetails() {
    logger := logger.GetLogger()
    logger.LogError("Payment failed.")
}
