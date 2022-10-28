package routers

import (
	"encoding/json" // permite gestionar la codificación JSON
	"net/http"      // permite gestionar la conectividad a través de WWW

	"github.com/jv2022/twittor-jvl/bd"     // package bd del proyecto
	"github.com/jv2022/twittor-jvl/models" // package models del proyecto
)

/*
ConsultaRelacion, cheque si existe relación entre 2 usuarios.
*/
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	// obtieniendo el parámetro ID del usuario que se está siguiendo
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	// cargando la información en el modelo Relacion
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	// verificando si existe la relación entre los 2 usuarios
	var resp models.RespuestaConsultaRelacion
	status, err := bd.ConsultoRelacion(t)

	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	// enviando la respuesta al HTTP
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)
}
