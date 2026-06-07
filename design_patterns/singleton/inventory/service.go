package inventory

import (
	"lld/design_patterns/singleton/logger"
)

type InventoryService struct{}

func (i *InventoryService) PrintInventoryDetails() {
    logger := logger.GetLogger()
    logger.LogInfo("Quantity is 1 for sku id 1.")
}