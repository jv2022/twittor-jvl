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
SubirBanner, sube el Banner al servidor, para un perfil determinado.
*/
func SubirBanner(w http.ResponseWriter, r *http.Request) {
	// obteniendo la información del banner a cargar, desde un formulario
	file, handle, err := r.FormFile("banner")

	// obtiendo la extensión del archivo que contiene el banner
	var extension = strings.Split(handle.Filename, ".")[1]

	// componiendo el nombre final del archivo donde se almacenará el banner
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	// creando el archivo que almacenará el banner
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	// copiando el archivo de banner desde el formulario a la carpeta en el servidor
	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	// registrando el nombre de archivo de banner en el perfil correspondiente
	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)

	if err != nil || status == false {
		http.Error(w, "Error al grabar el banner en la BD  ! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
