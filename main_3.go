package main

import "fmt"

var count int

var done = make(chan bool, 1)

func sayDoneNew(index int) {
	done <- true
	fmt.Printf("go %d done\n", index)
}
func waitUntilAllDoneNew(done chan bool, count int) {
	for count > 0 {
		if <-done {
			count--
		}
	}
}
func main() {
	fmt.Println("start....")
	fmt.Println("hello world")
	for i := 0; i < 5; i++ {
		count++

		go func(index int) {
			defer sayDoneNew(index)
			fmt.Printf("go %d running\n", index)
		}(i)
	}

	waitUntilAllDoneNew(done, count)
	fmt.Println("end....")
}
