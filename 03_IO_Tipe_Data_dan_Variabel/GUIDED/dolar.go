package main

import "fmt"

func konversiIDRtoUSD(idr int) float64 {
    kurs := 15000.0
    return float64(idr) / kurs
}


func main() {
    var idr int

    fmt.Print("Masukkan jumlah uang dalam IDR: ")
    fmt.Scan(&idr)

    usd := konversiIDRtoUSD(idr)

    fmt.Printf("Jumlah %d IDR dalam USD adalah %.2f\n", idr, usd)
}