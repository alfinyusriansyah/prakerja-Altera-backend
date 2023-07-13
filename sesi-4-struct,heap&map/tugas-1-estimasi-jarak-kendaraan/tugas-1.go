package main

import "fmt"

type Car struct {
	Tipe   string
	Fuelin float64
}

func Estimasi_jarak(c Car) float64 {
	bahan_bakar := 1.5
	estimasi_jarak := c.Fuelin * bahan_bakar
	return estimasi_jarak
}

func main() {
	car := Car{
		Tipe:   "Sedan",
		Fuelin: 10,
	}

	estimatedDistance := Estimasi_jarak(car)
	fmt.Printf("Perkiraan jarak yang bisa ditempuh dengan %s: %.2f mil\n", car.Tipe, estimatedDistance)
}
