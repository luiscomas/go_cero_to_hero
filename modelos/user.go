package modelos

import "time"

// IMPORTANTE: la carpeta modelos o Models es donde se definen las estructuras de datos que se van a utilizar en la aplicación.
// asi como tambien las funciones relacionadas con esas estructuras de datos.

type User struct {
	Id        int       `json:"id"` //para que el json lo tome en cuenta
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Status    bool      `json:"status"`
}

func (Usuario *User) AddUser(id int, username string, password string, email string, createdAt time.Time, status bool) User {
	Usuario.Id = id
	Usuario.Username = username
	Usuario.Password = password
	Usuario.Email = email
	Usuario.CreatedAt = createdAt
	Usuario.Status = status
	return User{id, username, password, email, createdAt, status}

}
