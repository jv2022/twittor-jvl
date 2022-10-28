package bd

import (
	"context" // permite acceder o establecer el contexto de ejecución de la aplicación
	"time"    // permite obtener el tiempo actual del sistema

	"github.com/jv2022/twittor-jvl/models" // package models del proyecto
)

/*
InsertoRelacion, graba la relación en la BD.
*/
func InsertoRelacion(t models.Relacion) (bool, error) {
	// creando un contexto asociado con un timeout de 15 segundos
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	// previniendo la cancelación del contexto, antes de finalizar la función
	defer cancel()

	// apuntando a la tabla (coleccion) relacion de la base de datos
	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	// insertando la relación en la bd
	_, err := col.InsertOne(ctx, t)

	if err != nil {
		return false, err
	}

	return true, nil
}
