package modelos

type Mujer struct {
	Hombre
}

func (m *Mujer) Respirar() {
	m.Respirando = true
}

func (m *Mujer) Comas() {
	m.Comiendo = true
}

func (m *Mujer) Pensar() {
	m.Pensando = true
}

func (m *Mujer) Sexo() string {
	return "mujer"
}
