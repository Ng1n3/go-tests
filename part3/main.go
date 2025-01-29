package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, error := json.Marshal(p1)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(string(bs))
}

/*
Start with this code. Instead of using the blank identifier, make sure the code is checking and
handling the error.
*/
