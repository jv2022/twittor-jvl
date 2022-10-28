package routers

import (
	"io"       // permite gestionar las operaciones de input/output
	"net/http" // permite conectarse a http
	"os"       // permite ejecutar instrucciones del sistema operativo
	"strings"  // permite gestionar cadenas

	"github.com/jv2022/twittor-jvl/bd"     // package bd del proyecto
	"github.com/jv2022/twittor-jvl/models" // package models del proyecto
)

/*
SubirAvatar, sube el Avatar al servidor, para un perfil determinado.
*/
func SubirAvatar(w http.ResponseWriter, r *http.Request) {
	// obteniendo la información del avatar a cargar, desde un formulario
	file, handle, err := r.FormFile("avatar")

	// obtiendo la extensión del archivo que contiene el avatar
	var extension = strings.Split(handle.Filename, ".")[1]

	// componiendo el nombre final del archivo donde se almacenará el avatar
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

	// creando el archivo que almacenará el avatar
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	// copiando el archivo de avatar desde el formulario a la carpeta en el servidor
	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	// registrando el nombre de archivo de avatar en el perfil correspondiente
	var usuario models.Usuario
	var status bool

	usuario.Avatar = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)

	if err != nil || status == false {
		http.Error(w, "Error al grabar el avatar en la BD  ! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
