package arreglos_slices

import "fmt"

var arreglo [10]int = [10]int{12, 14, 16, 18, 20, 22, 24, 26, 28, 30}
var tabla []int = []int{2, 4, 6}

func Slices() {
	porcion := arreglo[3:] //slice creado desde un vector, desde la posicion 3 hasta el final
	fmt.Println(porcion)
	fmt.Println(tabla)
	porcion2 := arreglo[:5]  //slice creado desde un vector, desde la posicion 0 hasta la 5
	porcion3 := arreglo[6:8] //slice creado desde un vector, desde la posicion 6 hasta la 8

	fmt.Println("porcion 1: Desde la posicion 3 hasta el final")
	fmt.Println(porcion)

	fmt.Println("porcion 2: Desde la posicion 0 hasta la 5")
	fmt.Println(porcion2)

	fmt.Println("porcion 3: Desde la posicion 6 hasta la 8")
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Println(elementos, "<-- Slice \nLargo de: ", len(elementos), "\nCapacidad de:", cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Println(nums, "<-- Slice \nLargo de: ", len(nums), "\nCapacidad de:", cap(nums))
}
