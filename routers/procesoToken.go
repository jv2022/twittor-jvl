package routers

import (
	"errors"  // permite gestionar los errores
	"strings" // permite gestionar las cadenas

	jwt "github.com/dgrijalva/jwt-go"      // json web token
	"github.com/jv2022/twittor-jvl/bd"     // package bd del proyecto
	"github.com/jv2022/twittor-jvl/models" // package models del proyecto
)

/*
Email, valor de Email usado en todos los EndPoints.
*/
var Email string

/*
IDUsuario, es el ID devuelto del modelo que se usará en todos los EndPonts.
*/
var IDUsuario string

/*
ProcesoToken, proceso token para extraer sus valores.
*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	// generando un slice con la clave a utilizar
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")

	// declarando una instancia en memoria de la estructura Claim
	claims := &models.Claim{}

	// filtrando la palabra "Bearer" del string que contiene el token
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	// limpiando los espacios blancos del token
	tk = strings.TrimSpace(splitToken[1])

	// parseando el token a la estructura Claim, haciendo uso de una función anónima
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	// verificando si se pudo parsear el token
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)

		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	// verificando si se el token era válido
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
