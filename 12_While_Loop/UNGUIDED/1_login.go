package main

import (
	"fmt"
)

func main() {
	const correctUsername = "Admin"
	const correctPassword = "Admin"

	var failedAttempts int
	var username, pass string

	for {
		fmt.Print("Masukkan username: ")
		fmt.Scanln(&username)

		fmt.Print("Masukkan password: ")
		fmt.Scanln(&pass)

		if username == correctUsername && pass == correctPassword {
			break
		} else {
			fmt.Println("Username atau password salah. Silakan coba lagi.")
			failedAttempts++
		}
	}

	fmt.Printf("Login berhasil! Percobaan gagal: %d kali\n", failedAttempts)
}