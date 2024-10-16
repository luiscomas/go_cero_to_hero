package ejercicios_interfaces

import (
	"fmt"

	"github.com/luiscomas/go_cero_to_hero/interfaces"
)

func HumanosRespirando(h interfaces.Humano) {
	h.Respirar()
	fmt.Printf("soy un %s, y estoy respirando\n", h.Sexo())
}
