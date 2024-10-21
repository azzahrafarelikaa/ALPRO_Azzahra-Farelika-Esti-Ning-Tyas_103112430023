package main

import (
	"fmt"
)

func main() {
	var jumlahBarang int

	// Pengguna menginputkan jumlah barang yang dibeli
	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scan(&jumlahBarang)

	totalPoin := 0

	// Menghitung poin
	for i := 1; i <= jumlahBarang; i++ {
		if i <= 5 {
			totalPoin += 10 // 10 poin untuk <+ 5 barang
		} else {
			totalPoin += 15 // 15 poin untuk > 5 barang
		}
	}

	// Output hasil menghitung poin belanja
	fmt.Printf("Total poin yang didapat: %d poin\n", totalPoin)
}