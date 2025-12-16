package main

import (
	"errors"
	"fmt"
)

func DoIt() (err error) {
	defer func() {
		p := recover()
		if p != nil { // 发生一个panic
			// 通过测试p值处理panic
			err = nil // 使包含函数不返回错误
		}
	}()

	// 任何可能引发panic的代码
	if err != nil {
		panic(errors.New(fmt.Sprintf("panic:%v", err)))
		// 或相当于 panic(fmt.Errorf("panic:%v", err))

	}

	return
}
func main() {
	DoIt()
}
