package paymentHandler

import (
	"github.com/taepunphu/go-sekai-shop-api/config"
	"github.com/taepunphu/go-sekai-shop-api/modules/payment/paymentUsecase"
)

type (
	PaymentQueueHandlerService interface {}

	paymentQueueHandler struct {
		cfg *config.Config
		paymentUsecase 	paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentQueueHandlerService(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentQueueHandlerService {
	return &paymentQueueHandler{
		cfg: cfg,
		paymentUsecase: paymentUsecase,
	}
}