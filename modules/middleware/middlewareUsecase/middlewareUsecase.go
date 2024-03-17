package middlewareUsecase

import "github.com/taepunphu/go-sekai-shop-api/modules/middleware/middlewareRepository"

type (
	MiddlewareUsecaseService interface {}

	middlewareUsecase struct {
		middlewareRepo middlewareRepository.MiddlewareRepositoryService
	}
)

func NewMiddlewareUsecaseService(middlewareRepo middlewareRepository.MiddlewareRepositoryService) MiddlewareUsecaseService {
	return &middlewareUsecase{middlewareRepo}
}

