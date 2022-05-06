package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

const OPENWEATHERMAP_BASE_URL = "https://api.openweathermap.org/data/2.5"
const OPENWEATHERMAP_API_KEY = "1e7e976727459533c8449ca9639e44e4"

type weather struct {
	temperature string `json:temperature`
	description string `json:description`
	humidity    string `json:humidity`
	windSpeed   string `json:windSpeed`
	city        string `json:city`
	country     string `json:country`
	icon        string `json:icon`
	error       string `json:error`
}

type weathers []weather

func Location(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	response, err := http.Get("http://ip-api.com/json/")

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(res, http.StatusText(http.StatusInternalServerError))
	} else {
		res.WriteHeader(http.StatusOK)
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(res, string(body))
	}
}

func CurrentByCity(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	params := mux.Vars(req)
	city := params["city"]

	response, err := http.Get(OPENWEATHERMAP_BASE_URL + "/weather?q=" + city + "&appid=" + OPENWEATHERMAP_API_KEY + "&units=metric")

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(res, http.StatusText(http.StatusInternalServerError))
	} else {
		res.WriteHeader(http.StatusOK)
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(res, string(body))
	}
}

func Current(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	city := "Cordoba"
	response, err := http.Get(OPENWEATHERMAP_BASE_URL + "/weather?q=" + city + "&appid=" + OPENWEATHERMAP_API_KEY + "&units=metric")

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(res, http.StatusText(http.StatusInternalServerError))
	} else {
		res.WriteHeader(http.StatusOK)
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(res, string(body))
	}
}

func ForecastByCity(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	params := mux.Vars(req)
	city := params["city"]
	response, err := http.Get(OPENWEATHERMAP_BASE_URL + "/forecast?q=" + city + "&appid=" + OPENWEATHERMAP_API_KEY + "&units=metric&cnt=5")

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(res, http.StatusText(http.StatusInternalServerError))
	} else {
		res.WriteHeader(http.StatusOK)
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(res, string(body))
	}
}

func Forecast(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	city := "Cordoba"
	response, err := http.Get(OPENWEATHERMAP_BASE_URL + "/forecast?q=" + city + "&appid=" + OPENWEATHERMAP_API_KEY + "&units=metric&cnt=5")

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(res, http.StatusText(http.StatusInternalServerError))
	} else {
		res.WriteHeader(http.StatusOK)
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(res, string(body))
	}
}
