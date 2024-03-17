package inventoryHandler

import (
	"github.com/taepunphu/go-sekai-shop-api/config"
	"github.com/taepunphu/go-sekai-shop-api/modules/inventory/inventoryUsecase"
)

type (
	InventoryHttpHandlerService interface {}

	inventoryHttpHandler struct {
		cfg *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryHttpHandlerService(cfg *config.Config, inventoryUsecase inventoryUsecase.InventoryUsecaseService) InventoryHttpHandlerService {
	return &inventoryHttpHandler{
		cfg: cfg,
		inventoryUsecase: inventoryUsecase,
	}
}