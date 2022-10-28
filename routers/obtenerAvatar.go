package routers

import (
	"io"       // permite gestionar las operaciones de input/output
	"net/http" // permite conectarse a http
	"os"       // permite ejecutar instrucciones del sistema operativo

	"github.com/jv2022/twittor-jvl/bd" // package bd del proyecto
)

/*
ObtenerAvatar, envia el Avatar al HTTP.
*/
func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {
	//obteniendo el ID del perfil
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	// buscando la información de perfil asociada al ID
	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	// abriendo el archivo de avatar asociada al perfil
	var OpenFile *os.File
	OpenFile, err = os.Open("uploads/avatars/" + perfil.Avatar)

	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	// enviando la imagen al HTTP
	_, err = io.Copy(w, OpenFile)

	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
	}
}
