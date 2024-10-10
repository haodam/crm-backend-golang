package usecase

import "github.com/haodam/user-backend-golang/internal/modules/user/repository"

type IUserService interface {
	Register(email string, purpose string) int
}
type userService struct {
	useRepository repository.IUserRepository
}

var _ IUserService = (*userService)(nil)

func NewUserService(useRepository repository.IUserRepository) IUserService {
	return &userService{
		useRepository: useRepository,
	}
}

func (us *userService) Register(email string, purpose string) int {
	return 0
}
