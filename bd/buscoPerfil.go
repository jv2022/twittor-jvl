package bd

import (
	"context" // permite gestionar el contexto de la aplicación
	"fmt"     // permite gestionar las entradas y salidas del terminal
	"time"    // permite gestionar el reloj interno

	"github.com/jv2022/twittor-jvl/models"       // package models del proyecto
	"go.mongodb.org/mongo-driver/bson"           // permite gestionar la base de datos
	"go.mongodb.org/mongo-driver/bson/primitive" // permite gestionar la base de datos
)

/*
BuscoPerfil, busca un perfil en la BD.
*/
func BuscoPerfil(ID string) (models.Usuario, error) {
	// creando el contexto de ejecución
	// estableciendo el defer para la liberación de recursos del contexto antes de finalizar la función
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// conectándose a la base de datos y a la tabla usuarios
	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	// formateando el ID de búsqueda
	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	// creando la condición de búsqueda para el ID
	condicion := bson.M{
		"_id": objID,
	}

	// buscando el perfil según el ID y blanqueando el campo password
	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""

	// verificanco si se encontró el perfil correspondiente al ID
	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return perfil, err
	}

	return perfil, nil
}
