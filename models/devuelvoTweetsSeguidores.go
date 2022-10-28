package models

import (
	"time" // permite gestionar el reloj interno

	"go.mongodb.org/mongo-driver/bson/primitive" // permite gestionar la base de datos MongoDB
)

/*
DevuelvoTweetsSeguidores, es la estructura con la que devolveremos los tweets de los seguidores.
*/
type DevuelvoTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID         string             `bson:"usuarioid" json:"userId,omitempty"`
	UsuarioRelacionID string             `bson:"usuariorealcionid" json:"userRelationId,omitempty"`

	Tweet struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
