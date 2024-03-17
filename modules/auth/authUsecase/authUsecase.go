package authUsecase

import "github.com/taepunphu/go-sekai-shop-api/modules/auth/authRepository"

type (
	AuthUsecaseService interface{}

	authUsecase struct {
		authRepo authRepository.AuthRepositoryService
	}
)

func NewAuthUsecaseService(authRepo authRepository.AuthRepositoryService) AuthUsecaseService {
	return &authUsecase{authRepo}
}