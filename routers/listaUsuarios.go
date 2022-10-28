package routers

import (
	"encoding/json" // permite gestionar la codificación JSON
	"net/http"      // permite gestionar la conexión HTTP
	"strconv"       // permite gestionar la conversión entre tipos de datos

	"github.com/jv2022/twittor-jvl/bd" // package bd del proyecto
)

/*
ListaUsuarios, leo la lista de los usuarios.
*/
func ListaUsuarios(w http.ResponseWriter, r *http.Request) {
	// rescatando los parámetros type, page, search
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "Debe enviar el parámero página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	// rescatando la información de los usuarios que seguimos o no seguimos
	pag := int64(pagTemp)
	result, status := bd.LeoUsuariosTodos(IDUsuario, pag, search, typeUser)

	if status == false {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(result)
}
