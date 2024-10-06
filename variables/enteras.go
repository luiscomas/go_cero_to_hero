package variables

import (
	"fmt"
)

func Variables_enteras() {
	fmt.Println("Variables enteras")
	//declaracion de variables enteras
	fmt.Println("Diversas formas de declarar variables enteras")
	a := 10
	var b int = 20
	c := int64(30)
	var d int32 = 40
	fmt.Println("a int = ", a)
	fmt.Println("b int = ", b)
	fmt.Println("c int64 = ", c)
	fmt.Printf("c int64 = %d\n", c)
	fmt.Println("d int32 = ", d)
}
