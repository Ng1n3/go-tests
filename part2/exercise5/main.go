package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

  go func() {
    c <- 89
  }()

	v, ok := <- c
	fmt.Println(v, ok)

	close(c)
	
	v, ok = <- c
	fmt.Println(v, ok)
}


/*
● Show the comma ok idiom starting with this code.
*/