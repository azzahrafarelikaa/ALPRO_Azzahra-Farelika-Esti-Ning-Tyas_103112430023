package main

import (
    "fmt"
)

type Karyawan struct {
    Nama      string
    GajiPokok float64
    Tunjangan float64
    Potongan  float64
}

func (k Karyawan) HitungTotalGaji() float64 {
    return k.GajiPokok + k.Tunjangan - k.Potongan
}

func main() {
    var karyawan Karyawan

    fmt.Print("Masukkan Nama Karyawan: ")
    fmt.Scanln(&karyawan.Nama)

    fmt.Print("Masukkan Gaji Pokok: Rp ")
    fmt.Scan(&karyawan.GajiPokok)

    fmt.Print("Masukkan Tunjangan: Rp ")
    fmt.Scan(&karyawan.Tunjangan)

    fmt.Print("Masukkan Potongan: Rp ")
    fmt.Scan(&karyawan.Potongan)

    totalGaji := karyawan.HitungTotalGaji()

    fmt.Println("\nInformasi Karyawan:")
    fmt.Println("Nama:", karyawan.Nama)
    fmt.Printf("Gaji Pokok: Rp %.2f\n", karyawan.GajiPokok)
    fmt.Printf("Tunjangan: Rp %.2f\n", karyawan.Tunjangan)
    fmt.Printf("Potongan: Rp %.2f\n", karyawan.Potongan)
    fmt.Printf("Total Gaji: Rp %.2f\n", totalGaji)
}
