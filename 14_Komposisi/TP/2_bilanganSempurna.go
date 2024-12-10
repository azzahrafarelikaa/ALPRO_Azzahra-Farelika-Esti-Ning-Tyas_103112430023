package main

import (
	"fmt"
)

func isPerfectNumber(num int) bool {
	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum == num
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	if isPerfectNumber(n) {
		fmt.Printf("Ya (karena faktor dari %d adalah ", n)
		for i := 1; i < n; i++ {
			if n%i == 0 {
				fmt.Print(i, " ")
			}
		}
		fmt.Printf("dan jumlahnya %d)\n", n)
	} else {
		fmt.Println("Tidak")
	}
}
