package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var incrementer int64
	wg.Add(30)

	for i := 0; i < 30; i++ {
		go func() {
      atomic.AddInt64(&incrementer, 1)
      fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", incrementer)
}

/*
Fix the race condition you created in exercise #4 by using package atomic
*/
