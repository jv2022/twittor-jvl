package bd

import "golang.org/x/crypto/bcrypt" // permite encriptar información

/*
EncriptarPassword, permite encriptar un password
*/
func EncriptarPassword(pass string) (string, error) {
	// costo de procesamiento para el algoritmo de encriptación a aplicar (a mayor número mayor es el costo)
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
