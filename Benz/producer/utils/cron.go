package utils

import (
	"time"

	"github.com/mithun/client/kafca"
	"github.com/mithun/service"
)

var (
	frequency time.Duration = 10 * time.Second
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func fuelLog(t time.Time) {

	var kafcaClient = kafca.GetNewKafca()
	var fuelService = service.GetNewFuelService(kafcaClient)

	err := fuelService.FuelLog("Bengaluru", "false", false)
	if err != nil {
		return
	}
}

//call fuelservice every 2 min (keep last parameter false)
func Schedulecal() {
	doEvery(frequency, fuelLog)
}
