package main

import "fmt"
func main() {
	var x int
	var y int
	var selesai bool

	fmt.Print("Masukan bilangan untuk cek kelipatan (contoh jika cek apakah 2 kelipatan dari 5-> 5 2):")
	fmt.Scan(&x, &y)
	for selesai = false; !selesai; {
		x = x - y
		fmt.Println(x)
		selesai = x <= 0
	}
	fmt.Println(x == 0)
}