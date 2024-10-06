package ejercicios

import (
	"strconv"
)

func Ejercicio01(valor string) (int, string) {

	parametro, err := strconv.Atoi(valor)
	if err != nil {
		return 0, "Hubo un error"
	}

	if parametro > 100 {
		return parametro, "Es mayor a 100"
	} else {
		return parametro, "Es menor a 100"
	}

}
