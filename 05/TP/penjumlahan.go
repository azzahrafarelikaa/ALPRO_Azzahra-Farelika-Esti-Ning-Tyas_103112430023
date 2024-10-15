package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel
	var n int64
	fmt.Print("Masukan angka n untuk dijumlah : ")
	// Input n
	fmt.Scan(&n)
	// Proses menghitung
	sum := (n * (n + 1) /2)
	// Output 
	fmt.Printf("Hasil jumlah semua bilangan urut dari 1 hingga %d adalah %d", n, sum)
}