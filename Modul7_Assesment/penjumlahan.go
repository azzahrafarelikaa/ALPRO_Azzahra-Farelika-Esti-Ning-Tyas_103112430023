package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	if x > y {
		fmt.Println("Nilai x harus lebih kecil atau sama dengan y!")
		return
	}

	sum := 0
	for i := x; i <= y; i++ {
		sum += i
	}

	fmt.Printf("Hasil penjumlahan dari %d sampai %d adalah: %d\n", x, y, sum)
}
