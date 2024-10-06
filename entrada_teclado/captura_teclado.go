package entrada_teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	fmt.Println("Ingrese el primer numero:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado no es un numero" + err.Error())
		}
	}
	fmt.Println("Ingrese el segundo numero:")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado no es un numero" + err.Error())
		}
	}
	fmt.Println("Ingrese el leyenda:")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)
}
