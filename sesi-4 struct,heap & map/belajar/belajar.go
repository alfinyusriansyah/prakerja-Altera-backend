package main

import "fmt"

type People struct {
	name    string
	age     int
	address string
}

func main() {
	var people = People{}
	people.name = "alfin"
	people.age = 23
	people.address = "genteng"

	tampilkanData(people)
}

func tampilkanData(people People) {
	fmt.Println(people)
}
