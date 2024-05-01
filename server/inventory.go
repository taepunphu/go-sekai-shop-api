package server

import (
	"github.com/taepunphu/go-sekai-shop-api/modules/inventory/inventoryHandler"
	"github.com/taepunphu/go-sekai-shop-api/modules/inventory/inventoryRepository"
	"github.com/taepunphu/go-sekai-shop-api/modules/inventory/inventoryUsecase"
)

func (s *server) inventoryService() {
	repo := inventoryRepository.NewInventoryRepositoryService(s.db)
	usecase := inventoryUsecase.NewInventoryUsecaseService(repo)
	httpHandler := inventoryHandler.NewInventoryHttpHandlerService(s.cfg, usecase)
	grpcHandler := inventoryHandler.NewInventoryGrpcHandlerService(usecase)
	queueHandler := inventoryHandler.NewInventoryQueueHandlerService(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventoryService := s.app.Group("/inventory_v1")

	//Health Check
	inventoryService.GET("", s.healthCheckService)
}