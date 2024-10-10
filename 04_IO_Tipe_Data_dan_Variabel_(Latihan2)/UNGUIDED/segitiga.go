package main

import (
	"fmt"
	"math"
)

func main() {
	var ax, ay, bx, by, cx, cy float64

	fmt.Println("Masukkan koordinat titik A (x y):")
	fmt.Scan(&ax, &ay)
	fmt.Println("Masukkan koordinat titik B (x y):")
	fmt.Scan(&bx, &by)
	fmt.Println("Masukkan koordinat titik C (x y):")
	fmt.Scan(&cx, &cy)
    
	var sisiAB, sisiBC, sisiCA, terpanjang float64
	sisiAB = math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	sisiBC = math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2)) 
	sisiCA = math.Sqrt(math.Pow(ax-cx, 2) + math.Pow(ay-cy, 2)) 

	terpanjang = sisiAB
	if sisiBC > terpanjang {
		terpanjang = sisiBC
	}
	if sisiCA > terpanjang {
		terpanjang = sisiCA
	}

	fmt.Printf("Panjang sisi terpanjang pada segitiga adalah: %.2f\n", terpanjang)
}