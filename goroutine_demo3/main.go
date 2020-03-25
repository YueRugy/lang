package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wait sync.WaitGroup
var once sync.Once

func worker(id int, jobs <-chan int, results chan<- int) {
	defer wait.Done()
	for v := range jobs {
		fmt.Printf("worker%d start job%d\n", id, v)
		time.Sleep(1 * time.Second)
		results <- v * 2
		fmt.Printf("worker%d end job%d", id, v)

	}
}

func produce(jobs chan<- int64) {
	defer close(jobs)
	defer wait.Done()
	for i := 0; i < 1000; i++ {
		jobs <- rand.Int63()
	}
}

func work(jobs <-chan int64, results chan<- int, tmp chan<- struct{}) {
	defer wait.Done()
	//defer once.Do(func() {
	//close(results)
	//})

	for {
		v, ok := <-jobs
		//fmt.Println(v)
		if !ok {
			tmp <- struct{}{}
			break
		}

		total := int64(0)
		for v > 0 {
			total += v % 10
			v = v / 10
		}
		results <- int(total)
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	tmp := make(chan struct{}, 24)
	jobs := make(chan int64, 1000)
	results := make(chan int, 1000)
	wait.Add(25)
	go produce(jobs)

	for i := 0; i < 24; i++ {
		go work(jobs, results, tmp)
	}

	for i := 0; i < 24; i++ {
		//v := <-tmp
		//if v == 1 && i == 23 {
		//	close(tmp)
		//	close(results)
		//}
		<-tmp
	}
	close(tmp)
	close(results)
	for v := range results {
		fmt.Println(v)
	}

	wait.Wait()

	//jobs := make(chan int, 100)
	//results := make(chan int, 100)
	//wait.Add(3)
	//for i := 1; i <= 3; i++ {
	//	go worker(i, jobs, results)
	//}

	//for i := 1; i <= 5; i++ {
	//	jobs <- i
	//}
	//close(jobs)
	//wait.Wait()
}
