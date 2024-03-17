package inventoryHandler

import (
	"github.com/taepunphu/go-sekai-shop-api/config"
	"github.com/taepunphu/go-sekai-shop-api/modules/inventory/inventoryUsecase"
)

type (
	InventoryQueueHandlerService interface {}

	inventoryQueueHandler struct {
		cfg *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryQueueHandlerService(cfg *config.Config, inventoryUsecase inventoryUsecase.InventoryUsecaseService) InventoryQueueHandlerService {
	return &inventoryQueueHandler{
		cfg: cfg,
		inventoryUsecase: inventoryUsecase,
	}
}