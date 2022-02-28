package main

import (
	"fmt"
	"os"
)

var NumeroPrimo = [6]int{2, 3, 5, 7, 11, 13}

var NumeroCompuesto = [6]int{4, 6, 8, 9, 10, 12}

func main() {

	fmt.Println(os.Args)

	if NumeroPrimo == NumeroCompuesto {
		fmt.Println("Es un numero primo")
	} else {
		fmt.Println("Es un numero compuesto")
	}
}
