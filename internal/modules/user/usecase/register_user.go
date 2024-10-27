package usecase

import (
	"context"
	"fmt"
	"github.com/haodam/user-backend-golang/internal/modules/user/repository"
	"github.com/haodam/user-backend-golang/utils/random"
	"time"
)

type IUserRegisterService interface {
	Execute(ctx context.Context, email string, purpose string) int
}

type userServiceUseCase struct {
	useRepository repository.IUserRegisterRepository
	otpRepository repository.IOtpRegisterRepository
}

var _ IUserRegisterService = (*userServiceUseCase)(nil)

func NewUserService(useRepository repository.IUserRegisterRepository, otpRepository repository.IOtpRegisterRepository) IUserRegisterService {
	return &userServiceUseCase{
		useRepository: useRepository,
		otpRepository: otpRepository,
	}
}

func (us *userServiceUseCase) Execute(ctx context.Context, email string, purpose string) int {

	// step 1: Hash email (ma hoa 1 chieu)

	// step 2: check email exists
	if us.useRepository.FindUserByEmail(ctx, email) {
		return 1
	}

	// step 3: new OTP
	otp := random.GenerateSixDigOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}
	fmt.Printf("OTP is ::: %d\n", otp)
	// step 4: save OTP in reids with expiration time (2 minute)
	err := us.otpRepository.GenOTP(email, otp, int64(10*time.Minute))
	if err != nil {
		fmt.Println(err)
	}

	// step 5: send email OTP
	// step 6: check OTP is available
	// step 7: check spam OTP
	return 0
}
