package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Menggunakan waktu sebagai seed untuk menghasilkan angka acak yang berbeda setiap kali program dijalankan
	rand.Seed(time.Now().UnixNano())

	// Program memilih angka acak antara 1 hingga 100
	randomNumber := rand.Intn(100) + 1

	// Variabel untuk menyimpan tebakan pengguna
	var guess int

	// Jumlah maksimal percobaan
	maxAttempts := 5

	fmt.Println("Selamat datang di permainan tebak angka!")
	fmt.Println("Saya telah memilih angka antara 1 hingga 100.")
	fmt.Printf("Anda memiliki %d kesempatan untuk menebak angka tersebut.\n", maxAttempts)

	// Loop untuk memberikan pengguna hingga 5 kali kesempatan
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		// Meminta input tebakan pengguna
		fmt.Printf("Percobaan %d: Masukkan tebakan Anda: ", attempts)
		fmt.Scan(&guess)

		// Mengecek apakah tebakan benar, terlalu besar, atau terlalu kecil
		if guess == randomNumber {
			fmt.Println("Selamat! Tebakan Anda benar!")
			return // Permainan berakhir jika tebakan benar
		} else if guess > randomNumber {
			fmt.Println("Terlalu besar!")
		} else {
			fmt.Println("Terlalu kecil!")
		}
	}

	// Jika pengguna tidak berhasil menebak angka dalam 5 percobaan
	fmt.Printf("Maaf, Anda telah menggunakan semua kesempatan. Angka yang benar adalah %d.\n", randomNumber)
}