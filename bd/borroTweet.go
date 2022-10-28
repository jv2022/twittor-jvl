package bd

import (
	"context" // permite gestionar el contexto de la aplicación
	"time"    // permite gestionar el reloj interno

	"go.mongodb.org/mongo-driver/bson"           // permite gestionar la base de datos MongoDB
	"go.mongodb.org/mongo-driver/bson/primitive" // permite gestionar la base de datos
)

/*
BorroTweet, borra un tweet de la base de datos, perteneciente a un perfil.
*/
func BorroTweet(ID string, UserID string) error {
	// creando el contexto de ejecución
	// estableciendo el defer para la liberación de recursos del contexto antes de finalizar la función
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// conectándose a la base de datos y a la tabla tweet
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	// formateando el ID de búsqueda del tweet
	objID, _ := primitive.ObjectIDFromHex(ID)

	// creando la condicion de busqueda
	condicion := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	// eliminando el tweet
	_, err := col.DeleteOne(ctx, condicion)
	return err
}
