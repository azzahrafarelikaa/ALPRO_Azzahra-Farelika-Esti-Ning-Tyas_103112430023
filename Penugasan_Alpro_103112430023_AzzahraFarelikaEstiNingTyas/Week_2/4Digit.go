package main

import (
    "fmt"
)

func main() {
    var x int
    fmt.Print("Masukkan bilangan (1-999): ")
    fmt.Scan(&x)

    d1 := x / 100
    d2 := (x / 10) % 10
    d3 := x % 10

    fmt.Printf("%d %d %d\n", d1, d2, d3)
}
