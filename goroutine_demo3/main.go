package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup

func worker(id int, jobs <-chan int, results chan<- int) {
	defer wait.Done()
	for v := range jobs {
		fmt.Printf("worker%d start job%d\n", id, v)
		time.Sleep(1 * time.Second)
		results <- v * 2
		fmt.Printf("worker%d end job%d\n", id, v)

	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	wait.Add(3)
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
	wait.Wait()
}
