package main

import (
	"fmt"
)

func main() {
	var x_0023 int

	fmt.Scan(&x_0023)

	if x_0023 < 2 {
		fmt.Println("False")
		return
	}

	prima := true
	for i := 2; i*i <= x_0023; i++ {
		if x_0023%i == 0 {
			prima = false
			break
		}
	}

	if prima {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
