package main

import (
	"fmt"
)

func main() {
	var harga, total int
	
	for {
		fmt.Print("Masukan harga barang (ketik 0 untuk selesai):")
		fmt.Scanln(&harga)

		if harga == 0 {
			break
		}

		total += harga
	}
	fmt.Printf("Total belanja Anda: %d\n", total)
}