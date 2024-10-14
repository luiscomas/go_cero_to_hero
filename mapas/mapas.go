package mapas

import "fmt"

func Mapas() {
	// miMapa := make(map[string]string)
	// fmt.Println(miMapa)

	// miMapa["Republica Dominicana"] = "Santo Domingo"
	// miMapa["Colombia"] = "Bogota"
	// fmt.Println(miMapa)
	// fmt.Println(miMapa["Colombia"])

	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 20,
		"Chivas":      37,
		"Cruz Azul":   36,
	}
	//usando la funcion range la cual actual como un foreach
	fmt.Println(campeonato)
	// for equipo, punntaje := range campeonato {
	// 	fmt.Println("Equipo:", equipo, "Puntos:", punntaje)
	// }
	//eliminar un elemento del mapa
	delete(campeonato, "Chivas")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Chivas"]
	fmt.Println(puntaje, existe)

	//agregar un elemento al mapa
	campeonato["River Plate"] = 35
	fmt.Println(campeonato)
}
