package routers

import (
	"net/http" // permite conectarse a http

	"github.com/jv2022/twittor-jvl/bd" // package bd del proyecto
)

/*
EliminarTweet, permite borrar un tweet de un perfil determinado.
*/
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	// obteniendo el ID del tweet a eliminar
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	// eliminando el tweet
	err := bd.BorroTweet(ID, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
