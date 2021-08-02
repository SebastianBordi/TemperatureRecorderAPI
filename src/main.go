package main

import (
	"log"

	rtr "github.com/sebastianbordi/TemperatureRegisterAPI/router"
	srv "github.com/sebastianbordi/TemperatureRegisterAPI/server"
)

const APP_VERSION = "0.1"
const PORT = "3456"
const BASE_URL = "/api/v1/"

func main() {
	log.Printf("Temperature and humidity register v %s", APP_VERSION)
	router := rtr.GetRouter(BASE_URL)
	server := srv.GetServer(PORT, &router)

	panic(server.ListenAndServe())
}
