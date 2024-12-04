package main

import (
	"fmt"
)

func main() {
	var targetDonasi int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scanln(&targetDonasi)

	var totalDonasi, jumlahDonatur int

	for totalDonasi < targetDonasi {
		var donasi int
		fmt.Print("Masukkan donasi dari donatur: ")
		fmt.Scanln(&donasi)

		totalDonasi += donasi
		jumlahDonatur++

		fmt.Printf("Total donasi sementara: %d\n", totalDonasi)
		fmt.Printf("Jumlah donatur sementara: %d\n", jumlahDonatur)
	}

	fmt.Printf("\nTarget donasi tercapai!\n")
	fmt.Printf("Total donasi: %d\n", totalDonasi)
	fmt.Printf("Jumlah donatur: %d\n", jumlahDonatur)
}
