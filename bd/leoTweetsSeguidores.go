package bd

import (
	"context" // permite gestionar el contexto de la aplicación
	"time"    // permite gestionar el reloj interno

	"github.com/jv2022/twittor-jvl/models" // package models del proyecto
	"go.mongodb.org/mongo-driver/bson"     // permite gestionar la base de datos MongoDB
)

/*
LeoTweetsSeguidores, lee los tweets de mis seguidores.
*/
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	// creando el contexto de ejecución
	// estableciendo el defer para la liberación de recursos del contexto antes de finalizar la función
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// conectándose a la base de datos y a la tabla relacion
	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	// calculando los registros a saltar antes de leer una página de inforamción
	skip := (pagina - 1) * 20

	// creando las condiciones necesarias para obtener los tweets de los seguidores
	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}}) // filtro
	condiciones = append(condiciones, bson.M{                                    // unión de tablas
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})          // genera la unión de tablas como un único detalle
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"fecha": -1}}) // sorteo descendente por fecha
	condiciones = append(condiciones, bson.M{"$skip": skip})                // configuración de salto
	condiciones = append(condiciones, bson.M{"$limit": 20})                 // máximo de registros a traer por página

	// generando el cursor para los tweets de todos los seguidores
	cursor, err := col.Aggregate(ctx, condiciones)

	// ejecutando el cursor (internamente decodifica los datos y los vuelca al slice)
	var result []models.DevuelvoTweetsSeguidores
	err = cursor.All(ctx, &result)

	if err != nil {
		return result, false
	}

	return result, true
}
