package modelos

// Se pueden usar en otros archivos solo las
// variables, estructuras y funciones que empiezan
// con mayusculas, sino, seran privadas

type Persona struct {
	IdPersona int `gorm:"primary_key"`
	Nombre    string
	DNI       string
}

func PruebaDeFuncion(valor string) {
	if valor == "hola" {
		println("Hola")
	}
}
