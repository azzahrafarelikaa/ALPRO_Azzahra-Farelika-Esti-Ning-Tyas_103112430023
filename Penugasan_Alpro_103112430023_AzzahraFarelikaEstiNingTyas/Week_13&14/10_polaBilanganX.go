package main

import "fmt"

func main() {
    var x_0023, baris, kolom int

    fmt.Scan(&x_0023)

    for baris = 1; baris <= x_0023; baris++ {
        for kolom = 1; kolom <= x_0023; kolom++ {
            if (baris == kolom) || (kolom+baris-1 == x_0023) {
                fmt.Print(baris, " ")
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }      
} 