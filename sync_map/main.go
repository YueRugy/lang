package main

import (
	"fmt"
	"strconv"
	"sync"
)

var rwLock sync.RWMutex

func set(m map[string]int, key string, v int) {
	defer rwLock.RUnlock()
	rwLock.RLock()
	m[key] = v
}

func get(m map[string]int, key string) int {
	defer rwLock.Unlock()
	rwLock.Lock()
	return m[key]
}

func main() {
	var wait sync.WaitGroup
	m := make(map[string]int, 20)

	for i := 0; i < 20; i++ {
		wait.Add(1)
		go func(value int) {
			defer wait.Done()
			key := strconv.Itoa(i)
			set(m, key, value)
			fmt.Println(get(m, key))
		}(i)
	}

	wait.Wait()
}
