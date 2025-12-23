package main

import (
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

var xlog = log.New(os.Stderr, "", log.Ltime+log.Lmicroseconds)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go SocketClientGo(&wg)

	ss := NewSocketServer()

	go func() {
		gid := getGID()
		err := ss.AcceptConnections(8080)
		if err != nil {
			xlog.Printf("%5d testSocketServer accept failed:%v\n", gid, err)
			return
		}
	}()

	wg.Wait()
	ss.Accepting = false
}

func SocketClientGo(wg *sync.WaitGroup) {
	defer wg.Done()
	gid := getGID()
	cmds := []string{TODCommand, SayingCommand}
	max := 10

	var xwg sync.WaitGroup
	for i := 0; i < max; i++ {
		xwg.Add(1)
		go func(index, max int) {
			defer xwg.Done()
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			sc := newSocketClient("127.0.0.1, 8080")
			xlog.Printf("%5d SocketClientGo request %d of %d\n", gid, index, max)
			resp, err := sc.GetCmd(cmds[rand.Intn(len(cmds))])
			if err != nil {
				xlog.Printf("%5d SocketClientGo failed :%v\n", gid, err)
				return
			}

			xlog.Printf("%5d SocketClientGo response %d of %d\n", gid, resp)
		}(i+1, max)
	}

	xwg.Wait()
}
