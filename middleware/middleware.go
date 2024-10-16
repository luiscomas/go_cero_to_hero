package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")

	result := operacionesMidd(sumar)(10, 5) //esta funcion operacionesmidd es una funcion middleware
	fmt.Println(result)

	result = operacionesMidd(restar)(100, 15)
	fmt.Println(result)

	result = operacionesMidd(multiplicar)(13, 50)

	fmt.Println(result)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operación")
		return f(a, b)
	}

}
