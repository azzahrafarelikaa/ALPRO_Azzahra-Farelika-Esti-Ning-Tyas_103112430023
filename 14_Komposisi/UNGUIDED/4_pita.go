package main

import "fmt"

func main() {
	var bunga, pita string
	var N int

	fmt.Print("N: ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)
		if i == N {
			pita += bunga
		} else {
			pita += bunga + "-"
		}
	}

	fmt.Println("Pita: ", pita)
}
