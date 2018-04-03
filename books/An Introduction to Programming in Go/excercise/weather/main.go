package main

import (
	"flag"
	"net/http"

	"./controllers"
	"./middlewares"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.MainController).Methods("GET")
	r.HandleFunc("/weather/{city}", controllers.WeaterByCityController).Methods("GET")
	r.Use(middlewares.LogingMiddleware)
	return r
}

func main() {
	var port string
	flag.StringVar(&port, "port", ":3000", "default port=:3000")
	flag.Parse()
	handler := cors.Default().Handler(NewRouter())
	http.ListenAndServe(port, handler)
}
