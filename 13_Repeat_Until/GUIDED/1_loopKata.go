package main

import "fmt"
func main() {
	var kata string
	var perulangan int

	fmt.Print("Masukan kata dan jumlah perulangan (contoh pagi 3):")
	fmt.Scan(&kata, &perulangan)
	counter := 0
	for done := false; !done; {
		fmt.Println(kata)
		counter++
		done = (counter >= perulangan)
	}
}