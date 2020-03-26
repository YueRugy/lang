package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64

var wait sync.WaitGroup
func add() {
	defer wait.Done()
	atomic.AddInt64(&x, 1)
}

func main() {
	for i:=0;i<100000;i++{
		wait.Add(1)
		go add()
	}
	//fmt.Println(x)
	wait.Wait()
	fmt.Println(x)
}
