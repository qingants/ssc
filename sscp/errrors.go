package sscp

import "fmt"

type Error struct {
	Code int
	Desc string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Desc)
}

// ErrIllegalMsg .
var ErrIllegalMsg = &Error{400, "illegal message"}
var ErrNotAcceptable = &Error{406, "Not Acceptable"}
