package bd

import (
	"context" // permite gestionar el contexto de la aplicación
	"fmt"
	"time" // permite gestionar el reloj interno

	"github.com/jv2022/twittor-jvl/models" // package models del proyecto
	"go.mongodb.org/mongo-driver/bson"     // permite gestionar la base de datos
)

/*
ConsultoRelacion, consulta la existencia de una relación entre 2 usuarios, en la base de datos.
*/
func ConsultoRelacion(t models.Relacion) (bool, error) {
	// creando el contexto de ejecución
	// estableciendo el defer para la liberación de recursos del contexto antes de finalizar la función
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// conectándose a la base de datos y a la tabla usuarios
	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	// creando la condición de búsqueda
	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	// buscando la relación determinada
	var resultado models.Relacion
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
