package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&number)

	if isPositive(number) {
		fmt.Printf("Bilangan %d adalah positif.\n", number)
	} else {
		fmt.Printf("Bilangan %d bukan positif.\n", number)
	}
}

func isPositive(n int) bool {
	return n > 0
}
