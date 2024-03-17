package paymentHandler

import (
	"github.com/taepunphu/go-sekai-shop-api/config"
	"github.com/taepunphu/go-sekai-shop-api/modules/payment/paymentUsecase"
)

type (
	PaymentHttpHandlerService interface {}

	paymentHttpHandler struct {
		cfg *config.Config
		paymentUsecase 	paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentHttpHandlerService(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentHttpHandlerService {
	return &paymentHttpHandler{
		cfg: cfg,
		paymentUsecase: paymentUsecase,
	}
}