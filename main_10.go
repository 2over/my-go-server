package main

import (
	"errors"
	"fmt"
)

type TryFunc func() error
type CatchFunc func(error) (rerr error, cerr error)
type FinallyFunc func()

type TryCatchError struct {
	tryError   error
	catchError error
}

func (tce *TryCatchError) Error() string {
	return tce.String()
}

func (tce *TryCatchError) String() string {
	return fmt.Sprintf("TryCatchError[%v %v]", tce.tryError, tce.catchError)
}

func (tce *TryCatchError) Cause() error {
	return tce.tryError
}
func (tce *TryCatchError) Catch() error {
	return tce.catchError
}

func TryFinally(t TryFunc, f FinallyFunc) (err error) {
	defer func() {
		f()
	}()

	err = t()
	if err != nil {
		err = &TryCatchError{err, nil}
	}

	return
}
func triageRecover(p interface{}, c CatchFunc) (err error) {
	if p != nil {
		var terr, cerr error
		if v, ok := p.(error); ok {
			terr = v
		}
		if xrerr, xcerr := c(terr); xrerr != nil {
			cerr = xcerr
			err = xrerr
		}

		if terr != nil || cerr != nil {
			err = &TryCatchError{terr, cerr}
		}
	}

	return err
}

func TryCatch(t TryFunc, c CatchFunc) (err error) {
	defer func() {
		if xerr := triageRecover(recover(), c); xerr != nil {
			err = xerr
		}
	}()

	err = t()
	return
}

func TryCatchFinally(t TryFunc, c CatchFunc, f FinallyFunc) (err error) {
	defer func() {
		f()
	}()

	defer func() {
		if xerr := triageRecover(recover(), c); xerr != nil {
			err = xerr
		}
	}()

	err = t()
	return
}

func main() {
	err := TryCatchFinally(func() error {
		fmt.Printf("in try\n")
		panic(errors.New("forced panic"))
	}, func(e error) (re, ce error) {
		fmt.Printf("in catch %v: %v %v\n", e, re, ce)
		return
	}, func() {
		fmt.Printf("in finally\n")

	})
	fmt.Printf("TCF returned: %v\n", err)
	err = TryFinally(func() error {
		fmt.Printf("in try\n")
		return errors.New("try err")
	}, func() {
		fmt.Printf("in finally\n")
	})
	fmt.Printf("TCF returned :%v\n", err)
	err = TryCatch(func() error {
		fmt.Printf("in try\n")
		panic(errors.New("forced panic"))
	}, func(e error) (re, ce error) {
		fmt.Printf("in catch %v: %v %v\n", e, re, ce)
		return
	})
	fmt.Printf("TCF returned: %v\n", err)
	err = TryCatch(func() error {
		fmt.Printf("in try\n")
		return nil
	}, func(e error) (re, ce error) {
		fmt.Printf("in catch %v: %v %v\n", e, re, ce)
		return
	})

	fmt.Printf("TCF returned: %v\n", err)
}
