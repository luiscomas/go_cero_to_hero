package variables

import (
	"fmt"
	"time"
)

var Variable_global = "Variable global"

const Constante_global = "Constante global"

func Global_y_constante() {
	fmt.Println("Variables globales y constantes")
	fmt.Println("Variable global = ", Variable_global)
	fmt.Println("Constante global = ", Constante_global)
}

func VariablesFloat() {
	fmt.Println("Variables float")
	//declaracion de variables float
	fmt.Println("Diversas formas de declarar variables float")
	a := 10.5
	var b float64 = 20.5
	c := float64(30.5)
	var d float32 = 40.5
	fmt.Println("a float = ", a)
	fmt.Println("b float = ", b)
	fmt.Println("c float64 = ", c)
	fmt.Printf("c float64 = %f\n", c)
	fmt.Println("d float32 = ", d)
}

func VariablesComplex() {
	fmt.Println("Variables complejas")
	//declaracion de variables complejas
	fmt.Println("Diversas formas de declarar variables complejas")
	a := 10 + 5i
	var b complex64 = 20 + 5i
	c := complex(30, 5)
	var d complex128 = 40 + 5i
	fmt.Println("a complex = ", a)
	fmt.Println("b complex = ", b)
	fmt.Println("c complex64 = ", c)
	fmt.Printf("c complex64 = %v\n", c)
	fmt.Println("d complex128 = ", d)
}

func VariablesString() {
	fmt.Println("Variables string")
	//declaracion de variables string
	fmt.Println("Diversas formas de declarar variables string")
	a := "Hola"
	var b string = "Mundo"
	c := `Hola "Mundo"`
	var d string = `Hola "Mundo"`
	fmt.Println("a string = ", a)
	fmt.Println("b string = ", b)
	fmt.Println("c string = ", c)
	fmt.Printf("c string = %s\n", c)
	fmt.Println("d string = ", d)
}

func VariablesRune() {
	fmt.Println("Variables rune")
	//declaracion de variables rune
	fmt.Println("Diversas formas de declarar variables rune")
	a := 'a'
	var b rune = 'b'
	c := rune('c')
	var d rune = 'd'
	fmt.Println("a rune = ", a)
	fmt.Println("b rune = ", b)
	fmt.Println("c rune = ", c)
	fmt.Printf("c rune = %c\n", c)
	fmt.Println("d rune = ", d)
}

func VariablesBoolean() {
	fmt.Println("Variables boolean")
	//declaracion de variables boolean
	fmt.Println("Diversas formas de declarar variables boolean")
	a := true
	var b bool = false
	c := bool(true)
	var d bool = false
	fmt.Println("a boolean = ", a)
	fmt.Println("b boolean = ", b)
	fmt.Println("c boolean = ", c)
	fmt.Printf("c boolean = %t\n", c)
	fmt.Println("d boolean = ", d)
}

func VariablesByte() {
	fmt.Println("Variables byte")
	//declaracion de variables byte
	fmt.Println("Diversas formas de declarar variables byte")
	a := byte(10)
	var b byte = 20
	c := byte(30)
	var d byte = 40
	fmt.Println("a byte = ", a)
	fmt.Println("b byte = ", b)
	fmt.Println("c byte = ", c)
	fmt.Printf("c byte = %d\n", c)
	fmt.Println("d byte = ", d)
}

func VariablesTime() {
	fmt.Println("Variables time")
	//declaracion de variables time
	fmt.Println("Diversas formas de declarar variables time")
	a := time.Now()
	var b time.Time = time.Now()
	fmt.Println("a time = ", a)
	fmt.Println("b time = ", b)
}
