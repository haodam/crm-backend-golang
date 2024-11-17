package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code,omitempty"`    // Status code
	Message string      `json:"message,omitempty"` // Thong bao loi
	Data    interface{} `json:"data,omitempty"`    // Tra ve du lieu
}

type ErrorResponseData struct {
	Code    int         `json:"code,omitempty"`    // Status code
	Message string      `json:"message,omitempty"` // Thong bao loi
	Detail  interface{} `json:"detail,omitempty"`  // Tra thong tin loi cu the
}

// Success response

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	if message == "" {
		message = msg[code]
	}
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
