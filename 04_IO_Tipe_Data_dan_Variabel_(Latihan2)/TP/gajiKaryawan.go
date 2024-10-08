package main

import (
	"fmt"
)

func main() {
	var jamKerja float64
	var upahPerJam float64
	const jamNormal = 40.0
	const mingguPerBulan = 4.0

	// Input jam kerja per minggu dan upah per jam
	fmt.Print("Masukkan jumlah jam kerja per minggu: ")
	fmt.Scan(&jamKerja)

	fmt.Print("Masukkan upah per jam: ")
	fmt.Scan(&upahPerJam)

	// Menghitung total gaji
	var gajiMingguan float64

	if jamKerja > jamNormal {
		jamLembur := jamKerja - jamNormal
		gajiMingguan = (jamNormal * upahPerJam) + (jamLembur * 1.5 * upahPerJam)
	} else {
		gajiMingguan = jamKerja * upahPerJam
	}

	// Menghitung gaji bulanan
	gajiBulanan := gajiMingguan * mingguPerBulan

	// Menampilkan total gaji bulanan
	fmt.Printf("Total gaji bulanan: Rp %.2f\n", gajiBulanan)
}
