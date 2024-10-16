package ejercicios_interfaces

import (
	"fmt"

	"github.com/luiscomas/go_cero_to_hero/interfaces"
)

func HumanosRespirando(h interfaces.Humano) {
	h.Respirar()
	h.EstaVivo()
	fmt.Printf("soy un %s, y estoy respirando\n", h.Sexo())
	fmt.Println("Estoy vivo? ", h.EstaVivo())
}
