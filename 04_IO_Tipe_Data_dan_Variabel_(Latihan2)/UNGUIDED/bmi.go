package main

import (
	"fmt"
)

func main () {
	var bmi, tinggi, berat float64 

	fmt.Print("Masukan nilai BMI:")
	fmt.Scan(&bmi)
	fmt.Print("Masukan tinggi badan (meter):")
	fmt.Scan(&tinggi)
	
	berat = bmi * (tinggi * tinggi)		// Menghitung bb
	fmt.Printf("Berat badan: %.2f kg\n", berat)
}