package middlew

import (
	"net/http" // permite conectarse a http

	"github.com/jv2022/twittor-jvl/bd" // package bd del proyecto
)

/*
ChequeoBD, middleware que permite conocer el estado de la base de datos.
*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConection() == 0 {
			http.Error(w, "Conexión perdidad con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
