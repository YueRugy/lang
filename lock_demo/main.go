package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x    int
	wait sync.WaitGroup
	lock sync.Mutex

	rwLock sync.RWMutex
)

func write() {
	defer wait.Done()
	//defer lock.Unlock()
	for i := 0; i < 50000; i++ {
		rwLock.Lock()
		//	lock.Lock()
		x += 1
		//	lock.Unlock()
		rwLock.Unlock()

	}
}

func read() {
	defer rwLock.RUnlock()
	//	defer lock.Unlock()
	defer wait.Done()
	rwLock.RLock()
	//lock.Lock()
	fmt.Println(x)

}

func main() {

	currentTime := time.Now()
	for i := 0; i < 2; i++ {
		wait.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wait.Add(1)
		go read()
	}

	//wait.Add(2)
	//go add()
	//go add()
	wait.Wait()

	fmt.Println(time.Now().Sub(currentTime))
	//fmt.Println(x)
}
