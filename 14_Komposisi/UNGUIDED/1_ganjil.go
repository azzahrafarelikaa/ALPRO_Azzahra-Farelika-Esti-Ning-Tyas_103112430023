package main

import "fmt"

func main() {
	var angka int

	fmt.Scanln(&angka)

	jumlahGanjil := 0
	for bilangan := 1; bilangan <= angka; bilangan++ {
		if bilangan%2 != 0 {
			jumlahGanjil++
		}
	}

	fmt.Printf("Terdapat %d bilangan ganjil", jumlahGanjil)
}
