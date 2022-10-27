package bd

import (
	"context" // permite gestionar el contexto de la aplicación
	"time"

	// permite gestionar las entradas y salidas del terminal
	"github.com/jv2022/twittor-jvl/models"       // package models del proyecto
	"go.mongodb.org/mongo-driver/bson"           // permite gestionar la base de datos
	"go.mongodb.org/mongo-driver/bson/primitive" // permite gestionar la base de datos
)

/*
ModificoRegistro, permite modificar el perfil de usuario.
*/
func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	// creando el contexto de ejecución
	// estableciendo el defer para la liberación de recursos del contexto antes de finalizar la función
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// conectándose a la base de datos y a la tabla usuarios
	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	// creando el mapa que almacenará los campos a modificar para el perfil
	registro := make(map[string]interface{})

	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}

	registro["fechaNacimiento"] = u.FechaNacimiento

	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}
	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}
	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}
	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}

	// creando la instrucción de actualización del perfil en la base de datos
	sentencia := bson.M{
		"$set": registro,
	}

	// creando el filtro para la actualización del perfil en la base de datos (_id == objID)
	objID, _ := primitive.ObjectIDFromHex(ID)
	//filtro := bson.M{"_id": bson.M{"$eq": objID}}
	filtro := bson.M{
		"_id": bson.M{
			"$eq": objID,
		},
	}

	// actualizando el perfil en la base de datos
	_, err := col.UpdateOne(ctx, filtro, sentencia)

	// verificando el resultado de la actualización
	if err != nil {
		return false, err
	}

	return true, nil
}
