package repository

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepositoryImpl struct{}

var _ IUserRepository = (*userRepositoryImpl)(nil)

func NewUserRepository() IUserRepository {
	return &userRepositoryImpl{}
}

func (u *userRepositoryImpl) GetUserByEmail(email string) bool {
	return false
}
