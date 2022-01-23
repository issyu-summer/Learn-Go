package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x int64
var lock sync.Mutex
var wg sync.WaitGroup

func add() {
	x++
	wg.Done()
}

func mutexAdd() {
	lock.Lock()
	x++
	lock.Unlock()
	wg.Done()
}

func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}
func main() {
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		//go add()
		//go mutexAdd()
		go atomicAdd()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x, end.Sub(start))
}
