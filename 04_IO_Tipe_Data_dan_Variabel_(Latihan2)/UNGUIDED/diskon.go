package main

import (
	"fmt"
)

func main() {
	var harga_awal, diskon, total int

	fmt.Print("Masukan total harga belanja:")
	fmt.Scan(&harga_awal)
	fmt.Print("Diskon:")
	fmt.Scan(&diskon)

	total = harga_awal - (harga_awal * (diskon / 100))

	fmt.Print("total harga:", total)
}