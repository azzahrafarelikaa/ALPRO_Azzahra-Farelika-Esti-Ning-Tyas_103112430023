package main

import (
	"fmt"
)

func main() {
	var transaksi_0023, saldo int

	saldo = 0

	for {
		fmt.Scan(&transaksi_0023)

		if transaksi_0023 == 0 {
			break
		}

		saldo += transaksi_0023
	}

	fmt.Print(saldo)
}
