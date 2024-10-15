package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel
	var n int
	// Input dari user
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&n)
	// Loop untuk mencetak segitiga bintang
	for i := 1; i <= n; i++ {
		// Loop untuk mencetak bintang sebanyak i pada setiap baris
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		// Pindah ke baris baru setelah mencetak bintang
		fmt.Println()
	}
}