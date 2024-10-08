package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	const phi = 3.14159

	// Input r
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&r)

	// Menghitung luas dan keliling
	luas := phi * math.Pow(r, 2)
	keliling := 2 * phi * r

	// Menampilkan hasil perhitungan
	fmt.Printf("Luas lingkaran: %.2f\n", luas)
	fmt.Printf("Keliling lingkaran: %.2f\n", keliling)
}
