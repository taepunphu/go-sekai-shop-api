package server

import (
	"github.com/taepunphu/go-sekai-shop-api/modules/payment/paymentHandler"
	"github.com/taepunphu/go-sekai-shop-api/modules/payment/paymentRepository"
	"github.com/taepunphu/go-sekai-shop-api/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	repo := paymentRepository.NewPaymentRepositoryService(s.db)
	usecase := paymentUsecase.NewPaymentUsecaseService(repo)
	httpHandler := paymentHandler.NewPaymentHttpHandlerService(s.cfg, usecase)
	queueHandler := paymentHandler.NewPaymentQueueHandlerService(s.cfg, usecase)

	_ = httpHandler
	_ = queueHandler

	payment := s.app.Group("/payment_v1")

	// Health Check
	payment.GET("", s.healthCheckService)
}