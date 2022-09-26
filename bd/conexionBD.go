package bd

import (
	"context" // permite acceder o establecer el contexto de ejecución de la aplicación
	"log"     // permite registrar información en el log de la aplicación

	"go.mongodb.org/mongo-driver/mongo"         // permite acceder a la base de datos MongoDB
	"go.mongodb.org/mongo-driver/mongo/options" // permite acceder a la base de datos MongoDB
)

/*
MongoCN, objeto de conexión a la base de datos
*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://jv2022:ayrsenna@twittor-jvl.8lxztis.mongodb.net/?retryWrites=true&w=majority")

/*
ConectarBD, función que permite conectar a la base de datos.
*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión Exitosa a la BD")
	return client
}

/*
ChequeoConection, función que realiza el ping a la base de datos.
*/
func ChequeoConection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1
}
