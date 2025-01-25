package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(30)
	counter := 0
	for i := 0; i < 30; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
      wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}

/*
Using goroutines, create an incrementer program
○ have a variable to hold the incrementer value
○ launch a bunch of goroutines
■ each goroutine should
Todd McLeod - Learn To Code Go - Page 134
● read the incrementer value
○ store it in a new variable
● yield the processor with runtime.Gosched()
● increment the new variable
● write the value in the new variable back to the incrementer
variable
● use waitgroups to wait for all of your goroutines to finish
● the above will create a race condition.
● Prove that it is a race condition by using the -race flag
● if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
*/
