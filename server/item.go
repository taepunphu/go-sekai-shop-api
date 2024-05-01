package server

import (
	"github.com/taepunphu/go-sekai-shop-api/modules/item/itemHandler"
	"github.com/taepunphu/go-sekai-shop-api/modules/item/itemRepository"
	"github.com/taepunphu/go-sekai-shop-api/modules/item/itemUsecase"
)

func (s *server) itemService() {
	repo := itemRepository.NewItemRepositoryService(s.db)
	usecase := itemUsecase.NewItemUsecaseService(repo)
	httpHandler := itemHandler.NewItemHttpHandlerService(s.cfg, usecase)
	grpcHandler := itemHandler.NewItemGrpcHandlerService(usecase)

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item_v1")

	// Health Check
	item.GET("", s.healthCheckService)
}