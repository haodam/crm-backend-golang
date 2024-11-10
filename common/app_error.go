package common

type Error struct {
	Message      string
	DebugMessage string
	Code         int
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Debug() string {
	return e.DebugMessage
}

func (e *Error) ErrCode() int {
	return e.Code
}
