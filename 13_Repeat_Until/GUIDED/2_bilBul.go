package main

import "fmt"
func main() {
	var angka int
	var selesai bool

	for selesai = true; selesai; {
		fmt.Scan(&angka)
		selesai = angka <= 0
	}
	fmt.Printf("%d adalah bilangan bulat positif\n", angka)
}