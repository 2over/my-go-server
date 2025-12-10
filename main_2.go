package main

import "fmt"

func main() {

	var count int
	// 在任何阻塞之前最多支持100
	var done = make(chan bool, 100)
	count++

	go func() {
		defer sayDone(done) // 必须是一个函数调用
	}()

	count++
	go func() {
		defer sayDone(done) // 必须是一个函数调用
	}()

	waitUntilAllDone(done, count)
	fmt.Println("done")

}

func sayDone(done chan bool) {
	fmt.Println("say done")
	done <- true
}

func waitUntilAllDone(done chan bool, count int) {
	for count > 0 {
		if <-done {
			count--
		}
	}
}
