package models

import (
	"time" // permite obtener el tiempo actual del sistema

	"go.mongodb.org/mongo-driver/bson/primitive" // permite modelar la información de entrada y salida a la base de datos
)

/*
DevuelvoTweets, es la estructura con la que devolveremos los Tweets.
*/
type DevuelvoTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userid" json:"userId,omitempty"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
