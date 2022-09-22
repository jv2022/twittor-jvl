package bd

import (
	"context" // permite acceder o establecer el contexto de ejecución de la aplicación
	"time"    // permite obtener el tiempo actual del sistema

	"github.com/jv2022/twittor-jvl/models" // package models del proyecto
	"go.mongodb.org/mongo-driver/bson"     // permite modelar la información de entrada y salida a la base de datos
)

/*
ChequeoYaExisteUsuario, recibe un email de parámetro y chequea si ya existe en la BD.
*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	// definiendo la condición de búsqueda para la BD
	condicion := bson.M{"email": email}

	var resultado models.Usuario
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
