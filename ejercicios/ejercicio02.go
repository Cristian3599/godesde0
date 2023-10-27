package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func multiplicacion(numero int) string {
	var texto string
	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)
	}

	return texto
}

func numeroValido(texto string) (bool, int) {
	numero, err := strconv.Atoi(texto)

	if err != nil {
		return false, 0
	}

	return true, numero
}

func IngresoNumerosMultiplicacion() string {
	scanner := bufio.NewScanner(os.Stdin)
	var numValido bool
	var numero int
	var texto string

	for {
		fmt.Println("Ingrese número: ")
		if scanner.Scan() {
			numValido, numero = numeroValido(scanner.Text())
			if numValido {
				break
			}
		}
	}

	texto = multiplicacion(numero)

	return texto
}
