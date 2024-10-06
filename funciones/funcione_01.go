package funciones

import (
	"strconv"
)

func Conviertoatexto(numero int) (string, bool) {
	var texto string
	texto = strconv.Itoa(numero)
	return texto, true
}
