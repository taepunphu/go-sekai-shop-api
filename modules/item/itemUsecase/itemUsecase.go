package itemUsecase

import (
	"github.com/taepunphu/go-sekai-shop-api/modules/item/itemRepository"
)

type (
	ItemUsecaseService interface {}

	itemUsecase struct {
		itemRepository itemRepository.ItemRepositoryService
	}
)

func NewItemUsecaseService(itemRepository itemRepository.ItemRepositoryService) ItemUsecaseService {
	return &itemUsecase{
		itemRepository: itemRepository,
	}
}