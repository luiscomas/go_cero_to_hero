package funciones

import "fmt"

func Exponencial(valor int) {
	if valor > 10000000 {
		return
	}
	fmt.Println(valor)
	Exponencial(valor * valor)
}
