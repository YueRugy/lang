package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var once sync.Once
var wait sync.WaitGroup

//var rand rand2.Rand

func f(ch chan int) {
	defer wait.Done()
	defer close(ch)

	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func f2(ch1, ch2 chan int) {

	defer wait.Done()
	defer once.Do(func() {
		close(ch2)
	})

	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
		fmt.Println(len(ch1))
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		//fmt.Println(rand.Intn(100))

	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	//	runtime.GOMAXPROCS(3)
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	wait.Add(3)
	go f(ch1)
	go f2(ch1, ch2)
	go f2(ch1, ch2)
	for v := range ch2 {
		fmt.Println(v)
		//if v==0{}
	}
	wait.Wait()
}
