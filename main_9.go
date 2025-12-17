package main

import (
	"errors"
	"fmt"
)

type ErrorWithCause struct {
	Err   error
	Cause error
}

func NewError(err error) *ErrorWithCause {
	return NewErrorWithCause(err, nil)
}

func NewErrorWithCause(err error, cause error) *ErrorWithCause {
	if err == nil {
		err = errors.New("no error supplied")
	}
	return &ErrorWithCause{err, cause}
}

func (wc ErrorWithCause) Error() string {
	xerr := wc.Err
	xcause := wc.Cause

	if xcause == nil {
		xcause = errors.New("no root cause supplied")
	}

	return fmt.Sprintf("ErrorWithCause{%v, %v}", xerr, xcause)
}

func (wc ErrorWithCause) String() string {
	return wc.Error()
}

func main() {
	//fmt.Printf("ErrorWithCause error:%s\n", ewc.Error())
	//fmt.Printf("ErrorWithCause value:%v\n\n", ewc)
}
