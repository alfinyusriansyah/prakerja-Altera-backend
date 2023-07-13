package main

import (
	"fmt"
)

func main() {
	// function
	hasil, pesan := tambah(2, 5)
	fmt.Println(hasil)
	fmt.Println(pesan)

	// slice
	slice := []int{2, 3, 4, 5}
	databar := []int{10}
	dataNilai := append(databar, slice...)
	fmt.Println(slice[:])
	fmt.Println(dataNilai)

	// map
	dataPen := map[string]int{}
	dataPen["januari"] = 1000
	dataPen["februari"] = 1000
	fmt.Println(dataPen)

	fmt.Println("==================================================")

}

func tambah(a, b int) (int, string) {
	luas := a * b
	return luas, "hola"
}
