package main

import "fmt"

// type People struct {
// 	name    string
// 	age     int
// 	address string
// }

// func main() {
// 	var people = People{}
// 	people.name = "alfin"
// 	people.age = 23
// 	people.address = "genteng"

// 	tampilkanData(people)
// }

// func tampilkanData(people People) {
// 	fmt.Println(people)
// }

type Student struct {
	Name   []string
	Scores []int
}

func main() {
	var student Student

	for i := 0; i < 5; i++ {
		fmt.Printf("input %v student name :", i+1)
		fmt.Scanf(student.Name[i])
	}
}
