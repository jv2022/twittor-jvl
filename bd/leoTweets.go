package bd

import (
	"context" // permite gestionar el contexto de la aplicación
	"log"     // permite gestionar el registro de eventos en el log de la aplicación
	"time"    // permite gestionar el reloj interno

	"github.com/jv2022/twittor-jvl/models"      // package models del proyecto
	"go.mongodb.org/mongo-driver/bson"          // permite gestionar la base de datos MongoDB
	"go.mongodb.org/mongo-driver/mongo/options" // permite configurar opciones de filtrado para la base de datos MongoDB
)

/*
LeoTweets, lee los tweets de un perfil de manera paginada.
*/
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	// creando el contexto de ejecución
	// estableciendo el defer para la liberación de recursos del contexto antes de finalizar la función
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// conectándose a la base de datos y a la tabla tweet
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	// creando el slice que retornará todos los tweets
	var resultados []*models.DevuelvoTweets

	// creando la condición para obtener los tweets del perfil seleccionado
	condicion := bson.M{
		"userid": ID,
	}

	// configurando las opciones de búsqueda para los tweets
	opciones := options.Find()
	opciones.SetLimit(20)                               // obtiene un máximo de 20 tweets
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) // ordena los tweets por fecha y descendentemente
	opciones.SetSkip((pagina - 1) * 20)                 // ubicación del puntero en la tabla tweet antes de obtener los tweets

	// buscando los tweets y creando el cursor para la paginación en la tabla de los tweets
	cursor, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	// recorriendo los tweets obtenidos y almacenándolos en el slice a retornar
	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)

		if err != nil {

			return resultados, false
		}

		resultados = append(resultados, &registro)
	}

	return resultados, true
}
