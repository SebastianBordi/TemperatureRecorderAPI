package view

import (
	"log"
	"net/http"

	contr "github.com/sebastianbordi/TemperatureRegisterAPI/controller"
	"github.com/sebastianbordi/TemperatureRegisterAPI/model"
)

type measurementResource struct {
	Controller *contr.MeasurementController
	Logger     *log.Logger
}

var measurementResourceInstance *measurementResource

func GetMasurementResource(controller *contr.Controller, logger *log.Logger) (*measurementResource, error) {
	if measurementResourceInstance == nil {
		measurementResourceInstance = &measurementResource{
			Controller: controller,
			Logger:     logger,
		}
	}
	return measurementResourceInstance, nil
}

func (resource *measurementResource) CreateMeasurement(rw http.ResponseWriter, r *http.Request) {

	measurement := &model.Measurement{}
	err := resource.Controller.Create(measurement)

}
