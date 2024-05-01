package server

import (
	"github.com/taepunphu/go-sekai-shop-api/modules/auth/authHandler"
	"github.com/taepunphu/go-sekai-shop-api/modules/auth/authRepository"
	"github.com/taepunphu/go-sekai-shop-api/modules/auth/authUsecase"
)


func (s *server) authService() {
	repo := authRepository.NewAuthRepositoryService(s.db)
	usecase := authUsecase.NewAuthUsecaseService(repo)
	httpHandler := authHandler.NewAuthHttpHandlerService(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandlerService(usecase)

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	// Health Check
	auth.GET("", s.healthCheckService)
}