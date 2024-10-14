package arreglos_slices

import "fmt"

type User string

func Arreglos() {
	fmt.Println("Arreglos")
	var arreglo [10]User
	//asignacion de valores
	arreglo[0] = "Juan"
	arreglo[1] = "Pedro"
	arreglo[7] = "Maria"

	var arreglo2 [5]int = [5]int{1, 2, 3, 4, 5}

	arreglo3 := [10]string{"Miguel", "jose", "Alberto"}
	// for i := 0; i < len(arreglo); i++ {
	// 	fmt.Scanf("ingrese un nombre: "+"%s", arreglo[i])
	// 	fmt.Println(arreglo[i] + " fue ingresado")
	// 	fmt.Println(arreglo)
	//}
	fmt.Println(arreglo)
	fmt.Println(arreglo2)
	fmt.Println(arreglo3)
	fmt.Println("**********************")
	for i := 0; i < len(arreglo3); i++ {
		fmt.Println(arreglo3[i] + "\n")
	}

	var matriz [5][4]int
	//asignacion de valores
	matriz[0][0] = 1

	fmt.Println(matriz)

}
