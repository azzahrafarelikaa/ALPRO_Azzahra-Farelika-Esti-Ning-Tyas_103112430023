package main

import (
	"fmt"
)

func main() {
	for selesai := false; !selesai; {
		var kanan, kiri float64
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kanan, &kiri)
		fmt.Println("Sepeda motor pak Andi akan oleng:", kiri-kanan >= 9)
		selesai = (kanan+kiri > 150) || (kanan < 0 || kiri < 0)
	}
	fmt.Println("Program selesai")
}