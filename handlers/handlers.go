package handlers

import (
	"log"      // permite registrar información en en log de la aplicación
	"net/http" // permite conectarse a http
	"os"       // permite acceder al sistema operativo

	"github.com/jv2022/twittor-jvl/middlew" // package middlew del proyecto
	"github.com/jv2022/twittor-jvl/routers" // package routers del proyecto

	"github.com/gorilla/mux" // permite manejar las peticiones http
	"github.com/rs/cors"     // permite dar permisos de ejecución remota a la aplicación
)

/*
Manejadores, función donde se setea el puerto, el handler y luego pone a escuchar al Servidor Web.
*/
func Manejadores() {
	router := mux.NewRouter()

	// definiendo los middleware a utilizar por cada endpoint de la aplicación
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")

	// obteniendo el puerto de conexión
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	// habilitando el acceso público de la aplicación
	handler := cors.AllowAll().Handler(router)
	// colocando la aplicación a escuchar en el Servidor + Puerto
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
