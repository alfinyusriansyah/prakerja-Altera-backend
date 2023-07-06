package main

import "fmt"

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	num := 5 // Ganti dengan bilangan yang ingin Anda cek

	for i := 0; i < num; i++ {
		if isPrime(i) {
			fmt.Printf("%d adalah bilangan prima.\n", i)
		} else {
			fmt.Printf("%d bukan bilangan prima.\n", i)
		}
	}

}
