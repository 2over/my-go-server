package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func LogWrapper(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		method, path := req.Method, req.URL
		fmt.Printf("entered handler for %s %s\n", method, path)
		f(w, req)
		fmt.Printf("exited handler for %s %s\n", method, path)
	}
}

func ElapsedTimeWrapper(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		method, path := req.Method, req.URL
		start := time.Now().UnixNano()
		f(w, req)
		fmt.Printf("elapsed time for %s %s:%ds\n", method, path, time.Now().UnixNano()-start)
	}
}

var specPort = ":8086"

func main() {
	// 常见HTTP请求处理程序
	handler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("in handler %v %v\n", req.Method, req.URL)
		time.Sleep(1 * time.Second)
		w.Write([]byte(fmt.Sprintf("In handler for %s %s\n", req.Method, req.URL)))
	}

	// 建议处理程序
	http.HandleFunc("/test", LogWrapper(ElapsedTimeWrapper(handler)))
	if err := http.ListenAndServe(specPort, nil); err != nil {
		log.Fatalf("Failed to start server on %s %v", specPort, err)
	}
}
