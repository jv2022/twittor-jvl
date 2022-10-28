package bd

import (
	"context" // permite gestionar el contexto de la aplicación
	"fmt"
	"time" // permite gestionar el reloj interno

	"github.com/jv2022/twittor-jvl/models"      // package models del proyecto
	"go.mongodb.org/mongo-driver/bson"          // permite gestionar la base de datos MongoDB
	"go.mongodb.org/mongo-driver/mongo/options" // permite configurar opciones de filtrado para la base de datos MongoDB
)

/*
LeoUsuariosTodos, lee los usuarios a los cuales se está o no está siguiendo.
*/
func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	// creando el contexto de ejecución
	// estableciendo el defer para la liberación de recursos del contexto antes de finalizar la función
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// conectándose a la base de datos y a la tabla usuarios
	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	// creando el slice que retornará los usuarios que seguimos
	var results []*models.Usuario

	// configurando las opciones de búsqueda para los usuarios que seguimos
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20) // ubicación del puntero en la tabla usuarios
	findOptions.SetLimit(20)             // devolverá un máximo de 20 usuarios

	// creando la condición para la búsqueda por nombre, de usuarios que seguimos
	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	// ejecutando la búsqueda
	cur, err := col.Find(ctx, query, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	// recorriendo los usuarios obtenidos en la búsqueda
	var encontrado, incluir bool

	for cur.Next(ctx) {
		// obtiendo la información para el modelo Usuarios
		var s models.Usuario
		err := cur.Decode(&s)

		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		// cargando la información en el modelo Relacion
		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		// verificando si el usuario del cursor, tiene relación con el usuario actual
		incluir = false
		encontrado, err = ConsultoRelacion(r)

		if tipo == "new" && encontrado == false {
			incluir = true // incluyendo usuarios que no estamos siguiendo
		}
		if tipo == "follow" && encontrado == true {
			incluir = true // incluyendo usuarios que si estamos siguiendo
		}
		if r.UsuarioRelacionID == ID {
			incluir = false // no incluyendo mi propio ID de usuario
		}

		// verificando si el usuario del cursor, debe incluirse en el slice a retornar
		if incluir == true {
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	// verificando si el cursor devolvió algún error
	err = cur.Err()

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	// cerrando el cursor
	cur.Close(ctx)

	return results, true
}
