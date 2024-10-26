package usecase

import (
	"context"
	"github.com/haodam/user-backend-golang/internal/modules/user/repository"
)

type IUserRegisterService interface {
	Execute(ctx context.Context, email string, purpose string) int
}

type userServiceUseCase struct {
	useRepository repository.IUserRegisterRepository
}

var _ IUserRegisterService = (*userServiceUseCase)(nil)

func NewUserService(useRepository repository.IUserRegisterRepository) IUserRegisterService {
	return &userServiceUseCase{
		useRepository: useRepository,
	}
}

func (us *userServiceUseCase) Execute(ctx context.Context, email string, purpose string) int {

	// step 1: Hash email (ma hoa 1 chieu)

	// step 2: check email exists
	if us.useRepository.GetUserByEmail(email) {
		return 1
	}

	// step 3: new OTP
	// step 4: save OTP in reids with expiration time (2 minute)
	// step 5: send email OTP
	// step 6: check OTP is available
	// step 7: check spam OTP
	return 0
}
