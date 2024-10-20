package main

import (
	"fmt"
)

func main (){
	var s, satu, dua int
	var hasil int
	
	fmt.Println("Masukan angka pertama untuk dikalikan :")
	fmt.Scan(&satu)
	fmt.Println("Masukan angka kedua untuk dikalikan :")
	fmt.Scan(&dua)
	hasil = 0
	for s = 1; s <= dua; s+=1 {
	hasil = hasil + satu
	}
	fmt.Println("Hasil perkalian dari dua bilangan tersebut adalah", hasil)
}