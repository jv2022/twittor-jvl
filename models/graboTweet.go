package models

import "time" // permite gestionar el reloj interno

/*
GraboTweet, tiene la información del tweet que se va a grabar en la base de datos.
*/
type GraboTweet struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
