package common

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestResponseErr(t *testing.T) {
	type args struct {
		c       *gin.Context
		code    int
		message []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ResponseErr(tt.args.c, tt.args.code, tt.args.message...)
		})
	}
}

func TestResponseErrs(t *testing.T) {
	type args struct {
		c       *gin.Context
		code    int
		err     error
		message []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ResponseErrs(tt.args.c, tt.args.code, tt.args.err, tt.args.message...)
		})
	}
}

func TestResponseOk(t *testing.T) {
	type args struct {
		c       *gin.Context
		code    int
		message string
		data    interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ResponseOk(tt.args.c, tt.args.code, tt.args.message, tt.args.data)
		})
	}
}

func TestSimpleResponseOK(t *testing.T) {
	type args struct {
		c    *gin.Context
		code int
		data interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SimpleResponseOK(tt.args.c, tt.args.code, tt.args.data)
		})
	}
}
