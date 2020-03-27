package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup

//var notify bool
var sli chan struct{}=make(chan struct{},1)

func main() {
	wait.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	//	notify=true
	sli <- struct{}{}
	defer close(sli)
	wait.Wait()

}
func f() {
	defer wait.Done()
FOR:
	for {
		fmt.Println("yue")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-sli:
			break FOR
		default:
		}

	}
}
