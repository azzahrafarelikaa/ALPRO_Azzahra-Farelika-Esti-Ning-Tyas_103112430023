package main

import (
	"fmt"
)

func main() {
	var kata string
	
	for {
		fmt.Print("Masukan kata:")
		fmt.Scanln(&kata)

		if kata == "telkom" {
			fmt.Println("Program selesai.")
			break
		} else {
			fmt.Println("Anda mengetik:", kata)
		}
	}
}