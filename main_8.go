package main

import (
	"errors"
	"fmt"
)

type MultiError []error

func (me MultiError) Error() (res string) {
	res = "MultiError"
	sep := " "
	for _, e := range me {
		res = fmt.Sprintf("%s%s%s", res, sep, e.Error())
		sep = ";"
	}

	return
}

func (me MultiError) String() string {
	return me.Error()
}

func main() {
	me := MultiError(make([]error, 0, 10))
	for _, v := range []string{"one", "two", "three"} {
		me = append(me, errors.New(v))
	}

	fmt.Printf("MultiError error:%s\n", me.Error())
	fmt.Printf("MultiError value:%v\n\n", me)
}
