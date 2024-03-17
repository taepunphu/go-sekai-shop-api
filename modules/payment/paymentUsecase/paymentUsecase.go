package paymentUsecase

import "github.com/taepunphu/go-sekai-shop-api/modules/payment/paymentRepository"

type (
	PaymentUsecaseService interface {}

	paymentUsecase struct {
		paymentRepository paymentRepository.PaymentRepositoryService
	}
)

func NewPaymentUsecaseService(paymentRepository paymentRepository.PaymentRepositoryService) PaymentUsecaseService {
	return &paymentUsecase{
		paymentRepository: paymentRepository,
	}
}