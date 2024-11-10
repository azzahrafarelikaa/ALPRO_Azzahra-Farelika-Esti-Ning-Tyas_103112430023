package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&number)

	absoluteValue := absolute(number)

	fmt.Printf("Nilai absolut dari %d adalah %d\n", number, absoluteValue)
}

func absolute(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
