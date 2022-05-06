package routes

import (
	"fmt"
	"net/http"

	"github.com/cattaneopablo/api-rest-go/controllers"
	"github.com/gorilla/mux"
)

func FunctionHandler(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "ROOT") })
	router.HandleFunc("/location", controllers.Location).Methods(http.MethodGet)
	router.HandleFunc("/current/{city}", controllers.CurrentByCity).Methods(http.MethodGet)
	router.HandleFunc("/current", controllers.Current).Methods(http.MethodGet)
	router.HandleFunc("/forecast/{city}", controllers.ForecastByCity).Methods(http.MethodGet)
	router.HandleFunc("/forecast", controllers.Forecast).Methods(http.MethodGet)
}
