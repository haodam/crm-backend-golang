package usecase

type IUserService interface {
}
type userService struct{}

var _ IUserService = (*userService)(nil)

func NewUserService() IUserService {
	return &userService{}
}
