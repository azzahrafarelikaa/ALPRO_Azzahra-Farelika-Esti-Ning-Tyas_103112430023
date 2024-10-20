package main

import (
	"fmt"
)

func main() {
	var a, b int
	var loop int

	fmt.Println("Masukan bilangan pertama:")
	fmt.Scan(&a)
	fmt.Println("Masukan bilangan terakhir:")
	fmt.Scan(&b)
	for loop = a; loop <=b; loop+=1 {
		fmt.Print(loop," ")
	}
}