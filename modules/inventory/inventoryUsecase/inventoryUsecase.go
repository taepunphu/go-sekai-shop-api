package inventoryUsecase

import "github.com/taepunphu/go-sekai-shop-api/modules/inventory/inventoryRepository"

type (
	InventoryUsecaseService interface {}

	inventoryUsecase struct {
		inventoryRepository inventoryRepository.InventoryRepositoryService
	}
)

func NewInventoryUsecaseService(inventoryRepository inventoryRepository.InventoryRepositoryService) InventoryUsecaseService {
	return &inventoryUsecase{
		inventoryRepository: inventoryRepository,
	}
}