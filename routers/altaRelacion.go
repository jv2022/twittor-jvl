package routers

import (
	"net/http" // permite conectarse a http

	"github.com/jv2022/twittor-jvl/bd"     // package bd del proyecto
	"github.com/jv2022/twittor-jvl/models" // package models del proyecto
)

/*
AltaRelacion, realiza el registro de la relación entre usuarios.
*/
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	//obtiendo el parámetro ID del usuario a seguir
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	// cargando la información para el modelo Relacion
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	// registrando la relación en la base de datos
	status, err := bd.InsertoRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
