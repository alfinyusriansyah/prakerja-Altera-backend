package main

import "fmt"

func luasTrapesium(sisi_a, sisi_b, tinggi float64) float64 {
	return (sisi_a + sisi_b) * tinggi / 2
}

func main() {
	sisi_a := 7.0
	sisi_b := 2.0
	tinggi := 4.0

	area := luasTrapesium(sisi_a, sisi_b, tinggi)

	fmt.Println("Luas trapesium :", "(", sisi_a, "+", sisi_b, ")", "x", tinggi, "x 1/2 =", area)
}
