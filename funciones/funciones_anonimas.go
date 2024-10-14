package funciones

import (
	"fmt"
)

func FuncionSuma() {
	suma := func(a int, b int) int {
		return a + b
	}

	fmt.Println("El resultado de su suma es :", suma(27, 05))
}
