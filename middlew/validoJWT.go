package middlew

import (
	"net/http" // permite conectarse a http

	"github.com/jv2022/twittor-jvl/routers" // package routers del proyecto
)

/*
ValidoJWT, permite validar el JWT que nos viene en la petición.
*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token !"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
