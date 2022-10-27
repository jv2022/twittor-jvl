package routers

import (
	"encoding/json" // permite codificar datos en json
	"net/http"      // permite conectarse a http
	"strconv"       // permite realizar conversiones entre tipos de datos

	"github.com/jv2022/twittor-jvl/bd" // package bd del proyecto
)

/*
LeoTweets, leo los tweets.
*/
func LeoTweets(w http.ResponseWriter, r *http.Request) {
	// obteniendo el ID del perfil del cual se obtendrán los tweets
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	// obteniendo el número de página de tweets a mostrar
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}

	// formateando el parámetro página a un entero
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "Debe enviar el parámetro página con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	// casteando el valor pagina a un int64
	pag := int64(pagina)
	respuesta, correcto := bd.LeoTweets(ID, pag)

	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(respuesta)
}
