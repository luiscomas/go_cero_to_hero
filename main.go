package main

import (
	"fmt"

	"github.com/luiscomas/go_cero_to_hero/goroutines"
	"github.com/luiscomas/go_cero_to_hero/webserver"
)

func main() {
	// variables.Variables_enteras()
	// variables.VariablesFloat()
	// variables.VariablesComplex()
	// variables.VariablesString()
	// variables.VariablesRune()
	// variables.VariablesBoolean()
	// //variables importadas de paquete time
	// variables.VariablesTime()

	// fmt.Println(variables.Variable_global)
	// fmt.Println(variables.Constante_global)

	// texto, verdad := funciones.Conviertoatexto(500)
	// fmt.Println(texto, verdad)

	//condiciones en GO
	//if
	// if os := runtime.GOOS; os == "Linux." || os == "Darwin." {
	// 	fmt.Println("Esto no es windows estas en ", os)
	// } else {
	// 	fmt.Println("Estas en", os)
	// }
	// fmt.Println("sentencia switch")
	// //switch
	// switch os := runtime.GOOS; os {
	// case "Linux":
	// 	fmt.Println("Esto no es windows estas en ", os)
	// case "Darwin":
	// 	fmt.Println("Esto no es windows estas en ", os)
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	// fmt.Println("resultado ejercicio # 1")
	// //ejercicio # 1
	// resultado, mensaje := ejercicios.Ejercicio01("500")
	// fmt.Println(resultado, mensaje)

	//Captura de teclado
	// entrada_teclado.IngresoNumeros()

	// pacquete iteraciones
	// iteraciones.Iterar()

	//ejercicio # 2
	// fmt.Println(ejercicio02.TablaMultiplicar())
	// files.SumaTabla()
	//files.LeoArchivo()
	// funciones.FuncionSuma()
	// funciones.Exponencial(10)
	//arreglos_slices.Arreglos()
	//arreglos_slices.Slices()
	//arreglos_slices.Capacidad()
	//mapas.Mapas()
	//users.AltaUsuario()

	//  interfaces
	// Pedro := new(modelos.Hombre)
	// e.HumanosRespirando(Pedro)

	// Maria := new(modelos.Mujer)
	// e.HumanosRespirando(Maria)

	//defer panic

	//d.MuestroDefer()
	//d.EjemploPanic()

	canal1 := make(chan bool)
	go goroutines.MiNombreLento("Luis", canal1)
	defer func() { <-canal1 }()
	fmt.Println("Estoy aquí")

	//canal2 := make(chan bool)

	webserver.WWebserver()

}
