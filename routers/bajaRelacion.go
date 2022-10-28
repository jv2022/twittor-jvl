package routers

import (
	"net/http" // permite conectarse a http

	"github.com/jv2022/twittor-jvl/bd"     // package bd del proyecto
	"github.com/jv2022/twittor-jvl/models" // package models del proyecto
)

/*
BajaRelacion, realiza el borrado de la relación entre usuarios.
*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	// rescatando el parámetro ID del usuario con el que se va a borrar la relación
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	// cargando la información al modelo Relacion
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	// eliminando la relación de la base de datos
	status, err := bd.BorroRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado borrar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
