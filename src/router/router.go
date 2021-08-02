package router

import (
	"github.com/gorilla/mux"
	v "github.com/sebastianbordi/TemperatureRegisterAPI/view"
)

func GetRouter(baseUrl string) mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(baseUrl+"measurement", v.CreateMeasurement)
	return *router
}
