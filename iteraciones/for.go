package iteraciones

import "fmt"

func Iterar() {
	//aqui veremos distintos tipos de for
	//for mas basico
	for {
		println("esto es un for basico infinito, detenido por el break")
		break
	}

	for i := 0; i < 10; i++ {
		fmt.Println("La iteracion de va por", i)
	}

	for i := 0; i < 100; i += 5 {
		fmt.Println("La iteracion de va por", i)
	}

	for i := 100; i > 0; i-- {
		fmt.Println("La iteracion de va por", i)
	}

	for i := 10; i > 1; i-- {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
