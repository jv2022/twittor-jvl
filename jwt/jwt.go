package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"      // JSON web token
	"github.com/jv2022/twittor-jvl/models" // package models del proyecto
)

/*
GeneroJWT, genera el encriptado con el JWT
*/
func GeneroJWT(t models.Usuario) (string, error) {
	// generando un slice con la clave a utilizar
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")

	// llenando el mapa que se utilizará para generar el JWT
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(24 * time.Hour).Unix(),
	}

	// generando el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	// firmando el token
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, err
}
