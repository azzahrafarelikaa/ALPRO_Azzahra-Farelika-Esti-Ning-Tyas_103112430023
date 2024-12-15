package main

import "fmt"

func main() {
	var angka int

	fmt.Scanln(&angka)
	
	if angka < 2 {
		fmt.Println("bukan prima")
		return
	}

	prima := true

	for pembagi := 2; pembagi*pembagi <= angka; pembagi++ {
		if angka%pembagi == 0 {
			prima = false
			break
		}
	}

	if prima {
		fmt.Println("prima")
	} else {
		fmt.Println("bukan prima")
	}
}
