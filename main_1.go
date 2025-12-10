package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("start")
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("finish")
		defer wg.Done() // 惯用的。 Done()相当于Add(-1)
	}()

	wg.Add(1)
	go func() {
		fmt.Println("finish")

		defer wg.Done()
	}()

	wg.Wait()
	fmt.Println("done")
}
