package main

import "fmt"

func main() {
	count := 100

	fmt.Println("bilangan kelipatan 7 dari 1 sampai ", count, ":")
	for i := 0; i < count; i++ {
		if i%7 == 0 {
			fmt.Print(i, ",")
		}
	}
}
