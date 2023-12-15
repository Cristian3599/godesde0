package main

import (
	"fmt"

	"github.com/godesde0/goroutines"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
	os := runtime.GOOS
	if os == "linux" {
		fmt.Println("Esto no es Windows, es ", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	numero, texto := ejercicios.ConvertiraEntero("69")
	fmt.Println(numero)
	fmt.Println(texto)*/

	//teclado.IngresoNumeros()

	//fmt.Println(ejercicios.IngresoNumerosMultiplicacion())

	//files.GrabaTabla()

	//files.SumaTabla()

	//files.LeoArchivo()

	//arreglosslices.Capacidad()

	//mapas.MostrarMapas()

	//users.AltaUsuario()

	/*Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)*/

	//d.EjemploPanic()

	canal1 := make(chan bool)
	go goroutines.MiNombreLentooo("Cristian Gomez", canal1)
	defer func() {
		<-canal1
	}()
	fmt.Println("Estoy aqui")

}
