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
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")
	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")

	// *** estos EndPoint no pueden ser probados en POSTMAN ***
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.ObtenerAvatar))).Methods("GET")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.ObtenerBanner))).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")
	router.HandleFunc("/listaUsuarios", middlew.ChequeoBD(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leoTweetsSeguidores", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweetsSeguidores))).Methods("GET")

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
