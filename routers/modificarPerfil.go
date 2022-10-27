package routers

import (
	"encoding/json" // permite gestionar la codificación JSON
	"net/http"      // permite gestionar la conectividad a través de WWW

	"github.com/jv2022/twittor-jvl/bd"     // package bd del proyecto
	"github.com/jv2022/twittor-jvl/models" // package models del proyecto
)

/*
ModificarPerfil, modifica el perfil de usuario.
*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	// decodificando la información del perfil recibida
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	// modificando el perfil en la base de datos
	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro. Reintente nuevamente"+err.Error(), 400)
		return
	}

	// verificando si se pudo realizar la actualización del perfil
	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario", 400)
		return
	}

	//
	w.WriteHeader(http.StatusCreated)
}
