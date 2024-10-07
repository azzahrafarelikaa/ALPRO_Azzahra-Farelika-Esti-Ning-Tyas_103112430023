package main

import "fmt"

func main() {
    var nama string
    var nim string
    var kelas string

    // Meminta input dari pengguna
    fmt.Print("Masukkan Nama: ")
    fmt.Scanln(&nama)

    fmt.Print("Masukkan NIM: ")
    fmt.Scanln(&nim)

    fmt.Print("Masukkan Kelas: ")
    fmt.Scanln(&kelas)

    // Menampilkan resume singkat
    fmt.Println("\nBiodata Mahasiswa:")
    fmt.Printf("Perkenalkan saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %s.\n", nama, kelas, nim)
}