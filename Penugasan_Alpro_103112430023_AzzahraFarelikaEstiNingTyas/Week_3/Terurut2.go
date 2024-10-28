package main

import (
    "fmt"
)

func main() {
    var x int
    fmt.Print("Masukkan bilangan 3 digit (100-999): ")
    fmt.Scan(&x)

    if x < 100 || x > 999 {
        fmt.Println("Masukan tidak valid. Harus berupa bilangan 3 digit.")
        return
    }
   
    d1 := x / 100
    d2 := (x / 10) % 10
    d3 := x % 10

    if d1 < d2 && d2 < d3 {
        fmt.Println("true - Urutan membesar")
    } else if d1 > d2 && d2 > d3 {
        fmt.Println("true - Urutan mengecil")
    } else {
        fmt.Println("false - Tidak terurut")
    }
}
