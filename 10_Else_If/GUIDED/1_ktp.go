package main

import "fmt" 

func main() {
	var usia int
	var kk bool 
	
	fmt.Print("Masukan umur anda:")
	fmt.Scan(&usia)
	fmt.Print("Sudah punya Kartu Keluarga (true/false):")
	fmt.Scan(&kk)

	if usia >= 17 && kk {
		fmt.Println("bisa membuat KTP") 
	}else{
		fmt.Println("belum bisa membuat KTP") }
	}