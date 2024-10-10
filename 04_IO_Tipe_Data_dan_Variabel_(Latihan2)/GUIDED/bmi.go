package main

import (
	"fmt"
)

func main() {
	var bb, tb, kg, hasil float64

	fmt.Print("Masukan berat badan:")
	fmt.Scan(&bb)
	fmt.Print("Masukan tinggi badan:")
	fmt.Scan(&kg)

	tb = kg * kg
	hasil = bb / tb

	fmt.Printf("Hasil Nilai BMI %.2f", hasil)
}