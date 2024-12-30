package main

import (
	"fmt"
)

func main() {
	var a_0023, b, c int

	fmt.Scan(&a_0023)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if a_0023+b > c && a_0023+c > b && b+c > a_0023 {
		if a_0023 == b && b == c {
			fmt.Println("Segitiga Sama Sisi")
		} else if a_0023 == b || a_0023 == c || b == c {
			fmt.Println("Segitiga Sama Kaki")
		} else {
			fmt.Println("Segitiga Sembarang")
		}
	}
}
