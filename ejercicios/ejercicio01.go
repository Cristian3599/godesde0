package ejercicios

import "strconv"

func ConvertiraEntero(texto string) (int, string) {
	numero, error := strconv.Atoi(texto)
	var mensaje string

	if error != nil {
		return 0, "Se presentó un error al convertir el texto a número"
	}

	if numero > 100 {
		mensaje = "Es mayor a 100"
	} else {
		mensaje = "Es menor a 100"
	}

	return numero, mensaje
}
