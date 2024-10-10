package main

import (
	"fmt"
)

func main() {
    // variabel
	var detik, jam, menit int
	fmt.Print("Masukan detik yang akan dikonversikan:")
	fmt.Scan(&detik)

	jam = detik / 3600
	menit = (detik % 3600) / 60
	detik = detik % 60
	
	fmt.Println(jam, "jam", menit, "menit dan", detik,"detik")
}