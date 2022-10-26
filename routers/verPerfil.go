package routers

import (
	"encoding/json" // permite gestionar la codificación JSON
	"net/http"      // permite gestionar la conectividad a través de WWW

	"github.com/jv2022/twittor-jvl/bd" // package bd del proyecto
)

/*
VerPerfil, permite extraer los valores del Perfil.
*/
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	// obteniendo el id asociado al perfil que se buscará
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	// buscando el perfil
	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar el registro "+err.Error(), 400)
		return
	}

	// retornando la información del perfil
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(perfil)
}
