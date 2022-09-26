package bd

import (
	"context" // permite acceder o establecer el contexto de ejecución de la aplicación
	"time"    // permite obtener el tiempo actual del sistema

	"github.com/jv2022/twittor-jvl/models"       // package models del proyecto
	"go.mongodb.org/mongo-driver/bson/primitive" // permite modelar la información de entrada y salida a la base de datos
)

/*
InsertoRegistro, es la parada final con la BD para insertar los datos del usuario
*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	// creando un contexto asociado con un timeout de 15 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// previniendo la cancelación del contexto, antes de finalizar la función
	defer cancel()

	// apuntando a la tabla (coleccion) usuarios de la base de datos
	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	// encriptando el password
	u.Password, _ = EncriptarPassword(u.Password)

	// insertando el usuario en la base de datos
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
