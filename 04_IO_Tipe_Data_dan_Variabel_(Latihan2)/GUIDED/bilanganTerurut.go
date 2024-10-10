package main

import (
	"fmt"
)

func main() {
	// deklarasi variabel
	var bilangan, d1, d2, d3 int

	fmt.Print("Masukan bilangan tiga digit:")
	fmt.Scan(&bilangan)

	d1 = bilangan / 100 // mengambil data pertama
	d2 = (bilangan % 100) / 10 // mengambil data kedua
	d3 = bilangan % 10 // mengambil data ketiga

	fmt.Println(d1 <= d2 && d2 <= d3) // cek digit
}