package repository

import "context"

type IUserRegisterRepository interface {
	GetUserByEmail(ctx context.Context, email string) bool
	FindUserByEmail(ctx context.Context, email string) bool
}

type userRepositoryImpl struct{}

var _ IUserRegisterRepository = (*userRepositoryImpl)(nil)

func NewUserRepository() IUserRegisterRepository {
	return &userRepositoryImpl{}
}

func (u *userRepositoryImpl) GetUserByEmail(ctx context.Context, email string) bool {
	return false
}

func (u *userRepositoryImpl) FindUserByEmail(ctx context.Context, email string) bool {
	return false
}
