package server

import (
	"github.com/taepunphu/go-sekai-shop-api/modules/player/playerHandler"
	"github.com/taepunphu/go-sekai-shop-api/modules/player/playerRepository"
	"github.com/taepunphu/go-sekai-shop-api/modules/player/playerUsecase"
)

func (s *server) playerService() {
	repo := playerRepository.NewPlayerRepositoryService(s.db)
	usecase := playerUsecase.NewPlayerUsecaseService(repo)
	httpHandler := playerHandler.NewPlayerHttpHandlerService(s.cfg, usecase)
	grpcHandler := playerHandler.NewPlayerGrpcHandlerService(usecase)
	queueHandler := playerHandler.NewPlayerQueueHandlerService(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/player_v1")

	// Health Checks
	player.GET("", s.healthCheckService)

}