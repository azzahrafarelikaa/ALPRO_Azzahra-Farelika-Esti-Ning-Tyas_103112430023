package main

import (
	"fmt"
)

func main() {
	var umur int
	var kk bool

	fmt.Print("Masukan umur anda :")
	fmt.Scan(&umur)
	fmt.Print("Sudah punya KK (true/false):")
	fmt.Scan(&kk)

	if umur >= 17 && kk {
		fmt.Print("bisa membuat KTP")
	} else {
		fmt.Print("belum bisa membuat KTP")
	}
}
