package routers

import (
	"encoding/json" // permite codificar datos en json
	"net/http"      // permite conectarse a http

	"github.com/jv2022/twittor-jvl/bd"     // package bd del proyecto
	"github.com/jv2022/twittor-jvl/models" // package models del proyecto
)

/*
Registro, función que permite crear en la BD el registro de usuario.
*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	// el Body es un objeto de tipo string, que se destruye una vez que es consumido
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el registro de usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
