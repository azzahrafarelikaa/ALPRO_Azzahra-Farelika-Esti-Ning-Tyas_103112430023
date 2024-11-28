package main

import "fmt"

func main() {
	const maxAttempts = 3
	var password string
	var correctPassword = "azzahra24"

	for i := 1; i <= maxAttempts; i++ {
		fmt.Print("Masukkan password:")
		fmt.Scanln(&password)

		if password == correctPassword {
			fmt.Println("Login berhasil")
			return
		} else {
			fmt.Println("Password salah, silahkan coba lagi.")
		}
	}
	fmt.Println("Login ditolak. Terlalu banyak percobaan gagal.")
}