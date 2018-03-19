package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func MainController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK!")
}

func WeaterByCityController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := vars["city"]
	fmt.Fprintln(w, "OK!", city)
}
