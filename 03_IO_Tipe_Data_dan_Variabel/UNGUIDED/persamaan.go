package main

import (
	"fmt"
)

func hitungX(fx float64) (float64, error) {
	// Jika fx = 5, maka error karena menyebabkan pembagian dengan 0
	if fx == 5 {
		return 0, fmt.Errorf("nilai f(x) tidak valid, karena menyebabkan pembagian dengan nol")
	}

	// Menghitung nilai x:
	// x = 2/(fx - 5) - 5
	x := 2/(fx-5) - 5
	return x, nil
}

func main() {
	var fx float64

	// Input untuk nilai x	fmt.Print("Masukkan nilai f(x): ")
	fmt.Scan(&fx)

	// Menghitung nilai x
	x, err := hitungX(fx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Nilai x yang dihitung adalah: %.4f\n", x)
	}
}