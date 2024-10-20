package main

import "fmt"

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Bilangan harus positif")
	} else {
		hasil := faktorial(n)
		fmt.Printf("Hasil faktorial dari %d adalah: %d\n", n, hasil)
	}
}