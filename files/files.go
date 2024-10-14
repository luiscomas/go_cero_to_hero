package files

import (
	"fmt"
	"os"

	"github.com/luiscomas/go_cero_to_hero/ejercicio02"
)

var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto = ejercicio02.TablaMultiplicar()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

func SumaTabla() {
	var texto = ejercicio02.TablaMultiplicar()
	if !append(fileName, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_RDONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el apend " + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo " + err.Error())
		return
	}
	//estp es una forma la mas comoda de leer un archivo
	fmt.Println(string(archivo))
	//scanner := bufio.NewScanner()
	// for scanner.Scan() {
	// 	registro := scanner.Text()
	// 	fmt.Println("> " + registro)
	// }
	// archivo.Close()
}
