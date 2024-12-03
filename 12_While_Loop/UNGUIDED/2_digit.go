package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&num)

	if num <= 0 {
		fmt.Println("Masukan harus bilangan bulat positif.")
		return
	}

	fmt.Println("Digit dari kanan ke kiri:")
	for num > 0 {
		digit := num % 10
		fmt.Println(digit)
		num /= 10
	}
}