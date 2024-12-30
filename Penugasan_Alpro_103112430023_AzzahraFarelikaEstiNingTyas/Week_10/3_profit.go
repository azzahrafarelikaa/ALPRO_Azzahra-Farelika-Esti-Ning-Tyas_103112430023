package main

import (
	"fmt"
)

func main() {
	var bulanIni_0023, bulanSebelumnya int

	fmt.Scan(&bulanIni_0023)
	fmt.Scan(&bulanSebelumnya)

	selisih := bulanIni_0023 - bulanSebelumnya

	if selisih > 0 {
		fmt.Print("Peningkatan sebesar ", selisih)
	} else if selisih < 0 {
		fmt.Print("Penurunan sebesar ", -selisih)
	} else {
		fmt.Println("Tetap")
	}
}
