package ejercicio02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() string {
	var numero int
	var err error
	//se agrego la parte de manejo de aerchivos aqui mismo
	var texto string

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese un número de tabla de multiplicar:")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				os.Clearenv()
			}

			for i := 0; i <= 10; i++ {
				//fmt.Println(numero, "x", i, "=", numero*i)
				texto += fmt.Sprintln(texto, numero, "x", i, "=", numero*i)
			}
			return texto

		}

	}

}
