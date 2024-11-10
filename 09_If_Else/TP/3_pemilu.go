package main

import (
	"fmt"
)

func main() {
	var umur int
	var negara string

	fmt.Print("Masukan umur anda :")
	fmt.Scan(&umur)
	fmt.Print("WNI atau WNA :")
	fmt.Scan(&negara)

	if umur >= 17 && negara == "WNI" {
		fmt.Print("Anda bisa mengikuti pemilu")
	} else {
		fmt.Print("Anda tidak bisa mengikuti pemilu")
	}
}
