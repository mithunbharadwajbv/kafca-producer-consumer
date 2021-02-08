package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mithun/models"
	"github.com/mithun/service"
)

type FuelController interface {
	FuelLog(*gin.Context)
}

type fuelController struct {
	service service.FuelService
}

func GetNewFuelController(service service.FuelService) FuelController {
	return &fuelController{service: service}
}

func (controller *fuelController) FuelLog(c *gin.Context) {
	var request models.LogFuelRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	err = controller.service.FuelLog(request.City, request.Flag, request.Mobile)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
}
