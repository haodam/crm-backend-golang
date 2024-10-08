package repository

type IUserRepository interface {
}

type userRepositoryImpl struct{}

func NewUserRepository() IUserRepository {
	return &userRepositoryImpl{}
}
