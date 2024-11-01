package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// OK { status: "OK_200", data: {} }
type OK struct {
	Status  string      `json:"status ,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Err { status: "ERR_500", message: "invalid request" }
type Err struct {
	Status  string   `json:"status ,omitempty"`
	Message string   `json:"message,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

func ResponseErr(c *gin.Context, code int, message ...string) {
	if len(message) == 0 {
		c.AbortWithStatusJSON(code, Err{
			Status: fmt.Sprintf("ERR_%d", code),
		})
		return
	}
	c.AbortWithStatusJSON(code, Err{
		Status:  fmt.Sprintf("ERR_%d", code),
		Message: message[0],
	})
}

func ResponseErrs(c *gin.Context, code int, err error, message ...string) {
	var resMsg string
	if len(message) == 0 {
		resMsg = http.StatusText(code)
	} else {
		resMsg = message[0]
	}
	var errs = strings.Split(err.Error(), "\n")
	c.AbortWithStatusJSON(code, Err{
		Status:  fmt.Sprintf("ERR_%d", code),
		Message: resMsg,
		Errors:  errs,
	})
}

func ResponseOk(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, OK{
		Status:  fmt.Sprintf("OK_%d", code),
		Message: message,
		Data:    data,
	})
}

func SimpleResponseOK(c *gin.Context, code int, data interface{}) {
	c.JSON(code, OK{
		Status:  fmt.Sprintf("OK_%d", code),
		Message: "OK",
		Data:    data,
	})
}
