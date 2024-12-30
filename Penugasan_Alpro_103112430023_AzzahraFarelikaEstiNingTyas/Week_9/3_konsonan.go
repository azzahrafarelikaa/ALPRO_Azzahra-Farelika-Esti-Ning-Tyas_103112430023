package main

import (
	"fmt"
)

func main() {
	var input_0023 string

	fmt.Scan(&input_0023)
	karakter := input_0023[0]

	if (karakter >= 'A' && karakter <= 'Z') || (karakter >= 'a' && karakter <= 'z') {
		switch karakter {
		case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
			fmt.Println("bukan konsonan")
		default:
			fmt.Println("konsonan")
		}
	} else {
		fmt.Println("bukan konsonan")
	}
}