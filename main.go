package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/cattaneopablo/api-rest-go/routes"
	"github.com/gorilla/mux"
)

const SERVER_PORT = 3000

func main() {
	router := mux.NewRouter().StrictSlash(true)
	routes.FunctionHandler(router)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(SERVER_PORT), router))
}
