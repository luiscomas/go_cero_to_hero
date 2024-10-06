package ejercicio02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() {
	var multiplicando int
	// var err error
	fmt.Println("Ingrese un número de tabla de multiplicar:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Debes ingresar un número")
		}
		multiplicando = numero
		for i := 0; i <= 10; i++ {
			fmt.Println(multiplicando, "x", i, "=", multiplicando*i)
		}

	}

}
