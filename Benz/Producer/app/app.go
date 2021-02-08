package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mithun/client/kafca"
	"github.com/mithun/controller"
	"github.com/mithun/service"
)

var (
	router = gin.Default()

	port = ":8080"
)

func StartApplication() {

	var kafcaClient = kafca.GetNewKafca()
	var fuelService = service.GetNewFuelService(kafcaClient)
	var fuelController = controller.GetNewFuelController(fuelService)

	router.POST("/logfuel", fuelController.FuelLog)
	router.Run(port)
}
