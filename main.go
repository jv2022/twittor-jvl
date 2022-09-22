package main

import (
	"log" // permite registrar información en el log de la aplicación

	"github.com/jv2022/twittor-jvl/bd"       // package bd del proyecto
	"github.com/jv2022/twittor-jvl/handlers" // package handlers del proyecto
)

func main() {
	if bd.ChequeoConection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}
