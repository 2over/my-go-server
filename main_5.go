package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

func CompressFileToNewGZIPFile(path string) (err error) {
	// 伪压缩码
	fmt.Printf("Starting compression of %s ...\n", path)
	start := time.Now()
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	end := time.Now()
	fmt.Printf("Compression of %s complete in %d seconds\n", path, end.Sub(start)/time.Second)
	return
}

func main() {
	var wg sync.WaitGroup
	for _, arg := range os.Args[1:] { // Args[0]是程序名
		wg.Add(1)

		go func(path string) {
			defer wg.Done()
			err := CompressFileToNewGZIPFile(path)
			if err != nil {
				log.Printf("File %s received error: %s\n", path, err)
				os.Exit(1)
			}
		}(arg) // 防止所有Go协程中的arg重复
	}

	wg.Wait()
}
