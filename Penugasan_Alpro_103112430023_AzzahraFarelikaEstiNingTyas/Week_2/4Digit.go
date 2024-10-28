package main

import (
    "fmt"
)

func main() {
    var x int
    fmt.Print("Masukkan bilangan (1-999): ")
    fmt.Scan(&x)

    if x < 1 || x > 999 {
        fmt.Println("Masukan tidak valid. Harus antara 1 hingga 999.")
        return
    }

    d1 := x / 100
    d2 := (x / 10) % 10
    d3 := x % 10

    fmt.Printf("%d %d %d\n", d1, d2, d3)
}
