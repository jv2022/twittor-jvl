package routers

import (
	"encoding/json" // permite gestionar la codificación JSON
	"net/http"      // permite gestionar la conexión HTTP
	"strconv"       // permite gestionar la conversión entre tipos de datos

	"github.com/jv2022/twittor-jvl/bd" // package bd del proyecto
)

/*
LeoTweetsSeguidores, lee los tweets de todos nuestros seguidores.
*/
func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {
	// obteniendo el párametro página
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero o mayor a 0", http.StatusBadRequest)
		return
	}

	// obteniendo todos los tweets de los seguidores
	respuesta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)

	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
