package main

import (
	"fmt"
)

func main() {
	var tebak_angka int
	kunci_jawaban := 1

	for {
		fmt.Print("Tebak angka (1-10):")
		fmt.Scan(&tebak_angka)

		if tebak_angka == kunci_jawaban {
			fmt.Println("Selamat, tebakan Anda benar!")
			break
		} else {
			fmt.Println("Tebakan Anda salah, coba lagi.")
		}
	}
}