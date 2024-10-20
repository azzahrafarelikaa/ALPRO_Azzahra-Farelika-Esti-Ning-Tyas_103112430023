package main

import (
	"fmt"
)

func main() {
	var s, alas, tinggi, n int
	var luas float64

	
	fmt.Println("Masukan nilai n :")
	fmt.Scan(&n)
	for s = 1; s <=n; s+=1 {
		fmt.Println("Masukan nilai alas dan tinggi segitiga :")
		fmt.Scan(&alas, &tinggi)
		luas = 0.5 * float64(alas * tinggi)
		fmt.Println(luas)
	}
}