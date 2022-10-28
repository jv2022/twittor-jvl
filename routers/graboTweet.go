package routers

import (
	"encoding/json" // permite codificar datos en json
	"net/http"      // permite conectarse a http
	"time"          // permite gestionar el reloj interno

	"github.com/jv2022/twittor-jvl/bd"     // package bd del proyecto
	"github.com/jv2022/twittor-jvl/models" // package models del proyecto
)

/*
GraboTweet, registra un tweet en la base de datos.
*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	// el Body es un objeto de tipo string, que se destruye una vez que es consumido
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	// formateando el registro del tweet a grabar en la base de datos
	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	//registrando el tweet en la base de datos
	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el Tweet", 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
