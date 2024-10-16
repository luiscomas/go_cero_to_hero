package defer_panic

import (
	"fmt"
	"log"
)

func MuestroDefer() {
	fmt.Println("Inicio de la funcion")
	defer fmt.Println("Fin de la funcion")
	fmt.Println("Esto esta dentro de la funcion")

}

//ejemplo panic

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error que genero panic \n %v", reco)
		}
	}() 
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
