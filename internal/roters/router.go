package roters

import "github.com/haodam/user-backend-golang/internal/roters/user"

type RouterGroup struct {
	User user.RouterGroupUser
	//Manager manager.Manager
}

var RouterGroupApp = new(RouterGroup)
