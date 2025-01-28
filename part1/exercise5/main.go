package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	// fmt.Println(users)

	// your code goes here

	fmt.Println("Sorted by Age: ")
	sort.Sort(ByAge(users))
	printUsers(users)

	// fmt.Println("--------------------------")
	fmt.Println("\nSorted by Last Name:")
	sort.Sort(ByLast(users))
	printUsers(users)
}


func printUsers(users []user) {
	for _, u := range users {
		fmt.Printf("Name: %s %s, Age: %d\n", u.First, u.Last, u.Age)
		fmt.Println("Sayings:")
		for _, saying := range u.Sayings {
			fmt.Printf("  - %s\n", saying)
		}
		fmt.Println()
	}
}

/*
Starting with this code, sort the []user by
● age
● last
Todd McLeod - Learn To Code Go - Page 129
Also sort each []string “Sayings” for each user
● print everything in a way that is pleasant
*/
