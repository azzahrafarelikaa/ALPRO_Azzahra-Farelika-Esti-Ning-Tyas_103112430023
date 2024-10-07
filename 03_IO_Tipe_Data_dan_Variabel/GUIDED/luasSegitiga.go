package main

import "fmt"

func hitungLuasSegitiga(alas, tinggi int) float64 {
    return 0.5 * float64(alas) * float64(tinggi)
}

func main() {
    var alas, tinggi int

    fmt.Print("Masukkan panjang alas segitiga: ")
    fmt.Scan(&alas)

    fmt.Print("Masukkan tinggi segitiga: ")
    fmt.Scan(&tinggi)

    luas := hitungLuasSegitiga(alas, tinggi)

    fmt.Printf("Luas segitiga dengan alas %d dan tinggi %d adalah %.2f\n", alas, tinggi, luas)
}