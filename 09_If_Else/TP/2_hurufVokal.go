package main

import "fmt"

func main() {
	var huruf string

	fmt.Print("Masukan huruf : ")
	fmt.Scan(&huruf)

	if huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" {
		fmt.Print("Huruf Vokal")
	} else {
		fmt.Print("Huruf Konsonan")
	}
}
