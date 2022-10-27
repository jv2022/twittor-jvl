package bd

import (
	"context" // permite gestionar el contexto de la aplicación
	"time"    // permite gestionar el reloj interno

	"github.com/jv2022/twittor-jvl/models"       // package models del proyecto
	"go.mongodb.org/mongo-driver/bson"           // permite gestionar la base de datos
	"go.mongodb.org/mongo-driver/bson/primitive" // permite modelar la información de entrada y salida a la base de datos
)

/*
InsertoTweet, graba el Tweet en la BD.
*/
func InsertoTweet(t models.GraboTweet) (string, bool, error) {
	// creando el contexto de ejecución
	// estableciendo el defer para la liberación de recursos del contexto antes de finalizar la función
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// conectándose a la base de datos y a la tabla tweet
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	// convirtiendo la información del registro de JSON a BSON
	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	// insertando el tweet en la base de datos
	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
