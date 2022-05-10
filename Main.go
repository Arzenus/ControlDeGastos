package main

import (
	modelos "main/Models" // Importa el paquete modelos,
	//la ruta tiene que ser el
	//NombreDelModulo/folderDondeEstaElPaquete
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost port=5432 user=postgres dbname=ControlGastos password=1 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	nuevoUsuario := modelos.Persona{Nombre: "Juan", DNI: "12345678"}

	if err != nil {
		panic(err)
	} else {
		db.Create(&nuevoUsuario) //Se pasa la referencia de la estructura
	}

	print(nuevoUsuario.IdPersona)
}
