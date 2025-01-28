package main

import (
	"fmt"
)

func main() {
	q := make(chan int, 1)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
    close(c)
    q <- 1
    close(q)
	}()

	return c
}

/*
Starting with this code, pull the values off the channel using a select statement
*/
