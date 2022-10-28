package bd

import (
	"context" // permite gestionar el contexto de la aplicación
	"time"    // permite gestionar el reloj interno

	"github.com/jv2022/twittor-jvl/models" // package models del proyecto
)

/*
BorroRelacion, borra la relación en la BD.
*/
func BorroRelacion(t models.Relacion) (bool, error) {
	// creando el contexto de ejecución
	// estableciendo el defer para la liberación de recursos del contexto antes de finalizar la función
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// conectándose a la base de datos y a la tabla relacion
	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	// eliminando la relación de la base de datos
	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}

	return true, nil
}
