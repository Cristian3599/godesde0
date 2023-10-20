package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func multiplicacion(numero int) {
	fmt.Println("----Multiplicación----")
	for i := 1; i <= 10; i++ {
		fmt.Println(numero * i)
	}
}

func numeroValido(texto string) (bool, int) {
	numero, err := strconv.Atoi(texto)

	if err != nil {
		return false, 0
	}

	return true, numero
}

func IngresoNumerosMultiplicacion() {
	scanner := bufio.NewScanner(os.Stdin)
	var numValido bool
	var numero int

	for {
		fmt.Println("Ingrese número: ")
		if scanner.Scan() {
			numValido, numero = numeroValido(scanner.Text())
			if numValido {
				break
			}
		}
	}

	multiplicacion(numero)
}
