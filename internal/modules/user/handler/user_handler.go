package handler

//type IUserHandler interface {
//	HandleUserRegister(ctx *gin.Context)
//	VerifyOTP(ctx *gin.Context)
//}
//
//type userHandlerImpl struct {
//	registerUserUseCase usecase.IUserRegister
//	verifyUserUseCase   usecase.IVerifyUserRegister
//}
//
//func NewUserHandler(d *database.Queries) IUserHandler {
//	return &userHandlerImpl{
//		registerUserUseCase: usecase.NewRegisterUserUseCase(d),
//		verifyUserUseCase:   usecase.NewVerifyUserUseCase(d),
//	}
//}
