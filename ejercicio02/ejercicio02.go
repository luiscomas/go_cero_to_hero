package ejercicio02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() {
	var numero int
	var err error

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
				fmt.Println(numero, "x", i, "=", numero*i)
			}

		}

	}

}
