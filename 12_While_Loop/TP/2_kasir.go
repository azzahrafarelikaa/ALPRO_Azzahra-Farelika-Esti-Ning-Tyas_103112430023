package main

import "fmt"

func main() {
	var totalHarga float64
	var namaBarang string
	var hargaBarang float64

	for {
		fmt.Print("Masukan nama barang (ketik 'selesai' untuk mengakhiri): ")
		fmt.Scanln(&namaBarang)
		if namaBarang == "selesai" {
			break
		}

		fmt.Print("Masukan harga barang: ")
		fmt.Scanln(&hargaBarang)

		totalHarga += hargaBarang

		fmt.Print("Barang", namaBarang, "berhasil ditambahkan.")
		fmt.Println("Total belanja saat ini: ", totalHarga)
	}
	fmt.Println("\nTotal belanja keseluruhan: ", totalHarga)
}