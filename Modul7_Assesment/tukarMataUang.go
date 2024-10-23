package main

import (
	"fmt"
)

func main() {
	var qirat int

	fmt.Print("Masukkan jumlah uang dalam qirat: ")
	fmt.Scan(&qirat)

	// Konversiii
	fals := qirat / 6
	sisaQirat := qirat % 6

	dirham := fals / 10
	sisaFals := fals % 10

	dinar := dirham / 10
	sisaDirham := dirham % 10

	fmt.Printf("Dinar: %d, Dirham: %d, Fals: %d, Qirat: %d\n", dinar, sisaDirham, sisaFals, sisaQirat)
}