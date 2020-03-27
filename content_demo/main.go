package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup
var notify bool

func main() {
	wait.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	notify=true
	wait.Wait()
}
func f() {
	defer wait.Done()
	for {
		fmt.Println("yue")
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}

	}
}
