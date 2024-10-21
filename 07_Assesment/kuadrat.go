package main

import (
	"fmt"
)

func main() {
	var n int

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan bulat positif n untuk menghitung hasil kuadrat: ")
	fmt.Scan(&n)

	// Output hasil perhitungan
	fmt.Println("Hasil kuadrat dari 1 hingga", n, "adalah :")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d^2 = %d\n", i, i*i)
	}
}