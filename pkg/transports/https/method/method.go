package method

type Method int

const (
	GET Method = iota + 1
	POST
	PUT
	PATCH
	DELETE
	HEAD
	OPTIONS
)
