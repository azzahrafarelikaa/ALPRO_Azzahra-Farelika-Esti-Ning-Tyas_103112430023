// a. Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal?
//Jawab:
//Jika user menginputkan nam bernilai 80.1 maka program akan mencetak BC karena sesuai dengan pernyataan kondisi if nam >= 72.5 yang menyatakan bahwa jika nam lebih dari sama dengan 72.5 maka variabel mk akan diatur ke BC.

// b. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya!
//Jawab:
//Urutan pernyataan kondisi dalam program tidak tersusun terurut dari terbesar ke terkecil, menyebabkan kesalahan dalam mengeksekusi program, program meletakan nam >= 72.5 setelah nam > 80 yang menyebabkan program memilih BC meskipun nam seharusnya dikategorikan sebagai AB.

// c.Perbaiki program tersebut!
// Jawab:
package main

import "fmt"

func main() {
    var nam float64
    var mk string

    fmt.Print("Nilai akhir mata kuliah: ")
    fmt.Scan(&nam)

    if nam >= 80 {
        mk = "A"
    } else if nam >= 72.5 {
        mk = "AB"
    } else if nam >= 65 {
        mk = "B"
    } else if nam >= 57.5 {
        mk = "BC"
    } else if nam >= 50 {
        mk = "C"
    } else if nam >= 40 {
        mk = "D"
    } else {
        mk = "E"
    }

    fmt.Println("Nilai mata kuliah: ", mk)
}