package main

import (
	"errors"
	"fmt"
	"time"
)

var NoError = errors.New("no error")

func GoroutineLauncher(gr func(), c *(chan error)) {
	go func() {
		defer func() {
			if p := recover(); p != nil {
				if c != nil {
					// 确保我们发送错误
					if err, ok := p.(error); ok {
						*c <- err
						return
					}

					*c <- errors.New(fmt.Sprintf("panic:%v", p))
				}

				return
			}
			if c != nil {
				*c <- NoError // 也可以发送nil并测试
			}
		}()

		gr()
	}()
}

var N = 5

func main() {
	var errchan = make(chan error, N) // N >= 1基于最大活动Go协程
	GoroutineLauncher(func() {
		time.Sleep(2 * time.Second)
		panic("panic happened!")
	}, &errchan)

	time.Sleep(5 * time.Second) // 模拟其他工作

	err := <-errchan
	if err != NoError {
		fmt.Printf("got %q", err.Error())
	}
}
