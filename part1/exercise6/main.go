package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("first go routine")
		wg.Done()
	}()
	go func() {
		fmt.Println("first go routine 2")
		wg.Done()
	}()

  fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())
}

/*
in addition to the main goroutine, launch two additional goroutines
○ each additional goroutine should print something out
● use waitgroups to make sure each goroutine finishes before your program exists
*/
