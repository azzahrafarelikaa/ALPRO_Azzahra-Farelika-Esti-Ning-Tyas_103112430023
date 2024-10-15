package main

import (
	"fmt"
)

func main() {
	// Loop dari 1 hingga 50
	for i := 1; i <= 50; i++ {
		// Mengecek apakah bilangan genap
		if i%2 == 0 {
			// Menampilkan bilangan genap
			fmt.Println(i)
		}
	}
}