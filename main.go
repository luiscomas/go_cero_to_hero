package main

import (
	"fmt"

	"github.com/luiscomas/go_cero_to_hero/funciones"
	"github.com/luiscomas/go_cero_to_hero/variables"
)

func main() {
	variables.Variables_enteras()
	// variables.VariablesFloat()
	// variables.VariablesComplex()
	// variables.VariablesString()
	// variables.VariablesRune()
	// variables.VariablesBoolean()
	// //variables importadas de paquete time
	// variables.VariablesTime()

	// fmt.Println(variables.Variable_global)
	// fmt.Println(variables.Constante_global)

	texto, verdad := funciones.Conviertoatexto(500)
	fmt.Println(texto, verdad)
	fmt.Println(texto)

}
