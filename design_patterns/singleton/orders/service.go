package orders

import (
	"lld/design_patterns/singleton/logger"
)

type OrderService struct{}

func (o *OrderService) PrintOrderDetails() {
    logger := logger.GetLogger()
    logger.LogInfo("order id is 1. Status is delivered.")
}
