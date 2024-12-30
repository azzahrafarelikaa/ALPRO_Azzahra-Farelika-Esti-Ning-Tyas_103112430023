package main

import (
	"fmt"
)

func main() {
	var shimmer_0023, shine int

	fmt.Scan(&shimmer_0023, &shine)

	if (shimmer_0023%2 == 0 && shine%2 != 0) || (shimmer_0023%2 != 0 && shine%2 == 0) {
		fmt.Println("berhasil")
	}
}
