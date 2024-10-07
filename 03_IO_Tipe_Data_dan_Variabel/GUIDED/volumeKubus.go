package main

import "fmt"

func hitungVolumeKubus(sisi int) int {
    return sisi * sisi * sisi
}

func main() {
    var sisi int

    fmt.Print("Masukkan panjang sisi kubus: ")
    fmt.Scan(&sisi)

    volume := hitungVolumeKubus(sisi)

    fmt.Printf("Volume kubus dengan sisi %d adalah %d\n", sisi, volume)
}