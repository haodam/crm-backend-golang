package initialize

import (
	"github.com/haodam/user-backend-golang/global"
	"github.com/haodam/user-backend-golang/internal/modules/user/repository"
	"github.com/haodam/user-backend-golang/internal/modules/user/usecase"
)

func ServiceInitialize() {
	queries := repository.New(global.MdbC)
	usecase.InitUserAuthed(usecase.NewAuthedUserUseCase(queries))
}
