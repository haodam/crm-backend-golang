package usecase

import (
	"context"
	"fmt"
	"github.com/haodam/user-backend-golang/internal/modules/user/entity"
	database "github.com/haodam/user-backend-golang/internal/modules/user/repository"
	"github.com/haodam/user-backend-golang/utils/crypto"
	"strings"
)

type registerUserUseCase struct {
	d *database.Queries
}

func newRegisterUserUseCase(d *database.Queries) *registerUserUseCase {
	return &registerUserUseCase{d: d}
}

//var _ IUserLogin = (*registerUserUseCase)(nil)

func (r registerUserUseCase) Register(ctx context.Context, req *entity.RegisterInput) (codeResult int, err error) {

	// Step1: Hash Email
	fmt.Printf("VerifyKey: %s\n", req.VerifyKey)
	fmt.Printf("VerifyType: %d\n", req.VerifyType)
	hashKey := crypto.GetHash(strings.ToLower(req.VerifyKey))
	fmt.Printf("hashKey: %s\n", hashKey)

	// Step2: Check user exists in uer base
	//userFound , err := r.d.

	// Step3: Create OTP

	// Step4: Generate OTP

	// Step5: Save OTP in Redis with expiration time

	// Step6: Send OTP

	// Step7: Save OTP to MYSQL

	// step8: gelatosID

	return 0, nil
}
