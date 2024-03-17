package middlewareRepository

type (
	MiddlewareRepositoryService interface {}

	middlewareRepository struct {}
)

func NewMiddlewareRepositoryService() MiddlewareRepositoryService {
	return &middlewareRepository{}
}