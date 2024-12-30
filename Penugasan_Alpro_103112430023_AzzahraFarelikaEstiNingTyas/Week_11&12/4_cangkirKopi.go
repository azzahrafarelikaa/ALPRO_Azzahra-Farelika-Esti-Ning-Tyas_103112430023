package main

import "fmt"

func main() {
	var n_0023, m, x, y int

	fmt.Scan(&n_0023, &m, &x, &y)

	var cups int

	for n_0023 >= x && m >= y {
		n_0023 -= x 
		m -= y 
		cups++ 
	}

	fmt.Print(cups)
}