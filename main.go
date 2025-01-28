// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	c := fanIn(boring("Joe"), boring("Ann"))
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(<-c)
// 	}
// 	fmt.Println("You're both boring; I'm leaving.")
// }

// func boring(msg string) <-chan string {
// 	c := make(chan string)
// 	go func() {
// 		for i := 0; ; i++ {
// 			c <- fmt.Sprintf("%s %d", msg, i)
// 			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
// 		}
// 	}()
// 	return c
// }

// // FAN IN
// func fanIn(input1, input2 <-chan string) <-chan string {
// 	c := make(chan string)
// 	go func() {
// 		for {
// 			c <- <-input1
// 		}
// 	}()
// 	go func() {
// 		for {
// 			c <- <-input2
// 		}
// 	}()
// 	return c
// }

/*
code source:
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25

source:
https://blog.golang.org/pipelines
*/



package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
