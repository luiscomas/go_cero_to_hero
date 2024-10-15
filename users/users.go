package users

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/luiscomas/go_cero_to_hero/modelos"
)

func AltaUsuario() {
	usuario := new(modelos.User)
	// usuario.Username = "luiscomas"
	// usuario.Password = "123456"
	// usuario.Email = "cronos@gmail.com"
	// usuario.CreatedAt = time.Now()
	// usuario.Status = true
	// usuario.Id = 1
	usuario.AddUser(1, "luiscomas", "123456", "cronos@gmail.com", time.Now(), true)
	covnertido, err := json.Marshal(usuario)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(covnertido))
	//fmt.Println(usuario)
}
