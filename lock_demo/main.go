package main

import (
	"fmt"
	"sync"
)

var (
	x    int
	wait sync.WaitGroup
	lock sync.Mutex
)

func add() {
	defer wait.Done()
	//defer lock.Unlock()
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

func main() {
	wait.Add(2)
	go add()
	go add()
	wait.Wait()

	fmt.Println(x)
}
