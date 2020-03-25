package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wait sync.WaitGroup

func f(i int) {
	defer wait.Done()
	fmt.Println(i)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
}

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i <= 100; i++ {
		wait.Add(1)
		f(i)
		//go hello(i)
		//		go func(i int ) {
		//			fmt.Println(i)
		//		}(i)
		//currentTime := time.Now()
		//go func(t time.Time) {
		//	fmt.Println(time.Now().Sub(currentTime))
		//}(currentTime)
		//go hello(i)

	}
	wait.Wait()
	fmt.Println("main")
	//time.Sleep(time.Second * 1)

}

func hello(i int) {
	fmt.Printf("hello%d\n", i)
}
