package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup

//var notify bool

func main() {
	wait.Add(1)
	ctx, cf := context.WithCancel(context.Background())
	go f(ctx)
	time.Sleep(time.Second * 5)
	cf()
	//notify=true
	wait.Wait()
}
func f(ctx context.Context) {
	defer wait.Done()
	go f2(ctx)
FOR:
	for {
		fmt.Println("yue")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FOR
		default:

		}
	}
}
func f2(ctx context.Context) {
	//defer wait.Done()
FOR:
	for {
		fmt.Println("wei")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FOR
		default:

		}
	}
}
