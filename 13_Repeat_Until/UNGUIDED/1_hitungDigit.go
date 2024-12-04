package main

import (
	"fmt"
)

func main() {
	var digit int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&digit)

	if digit <= 0 {
		fmt.Println("Bilangan harus bulat positif.")
		return
	}

	count := 0
	for digit > 0 {
		digit /= 10
		count++
	}

	fmt.Printf("Banyaknya digit: %d\n", count)
}
