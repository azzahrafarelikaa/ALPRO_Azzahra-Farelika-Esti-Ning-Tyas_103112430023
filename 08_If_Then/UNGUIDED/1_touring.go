package main

import (
	"fmt"
)

func hitungMotor(tOrang int) int {
	return (tOrang + 1) / 2
}

func main() {
	var tOrang int
	fmt.Print("Masukkan jumlah orang yang mengikuti touring: ")
	fmt.Scan(&tOrang)

	tMotor := hitungMotor(tOrang)
	fmt.Printf("Jumlah motor yang dibutuhkan: %d\n", tMotor)
}
