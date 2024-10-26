package repository

type IUserRegisterRepository interface {
	GetUserByEmail(email string) bool
}

type userRepositoryImpl struct{}

var _ IUserRegisterRepository = (*userRepositoryImpl)(nil)

func NewUserRepository() IUserRegisterRepository {
	return &userRepositoryImpl{}
}

func (u *userRepositoryImpl) GetUserByEmail(email string) bool {
	return false
}
