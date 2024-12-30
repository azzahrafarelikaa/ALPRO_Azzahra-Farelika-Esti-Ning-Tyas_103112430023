package main

import (
	"fmt"
)

func main() {
	var tim1_0023, tim2, tim3, tim4 int

	fmt.Scan(&tim1_0023)
	fmt.Scan(&tim2)
	fmt.Scan(&tim3)
	fmt.Scan(&tim4)

	terbanyak := tim1_0023
	if tim2 > terbanyak {
		terbanyak = tim2
	}
	if tim3 > terbanyak {
		terbanyak = tim3
	}
	if tim4 > terbanyak {
		terbanyak = tim4
	}

	tersedikit := tim1_0023
	if tim2 < tersedikit {
		tersedikit = tim2
	}
	if tim3 < tersedikit {
		tersedikit = tim3
	}
	if tim4 < tersedikit {
		tersedikit = tim4
	}

	fmt.Print(terbanyak, tersedikit)
}
