package models

import (
	jwt "github.com/dgrijalva/jwt-go"            // json web token
	"go.mongodb.org/mongo-driver/bson/primitive" // acceso a la bd mongodb
)

/*
Claim, es la estructura utilizada para procesar el JWT.
*/
type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`

	// heredando la estructura jwt.StandarClaims
	jwt.StandardClaims
}
