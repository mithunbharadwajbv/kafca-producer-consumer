package service

import (
	"errors"
	"fmt"

	"github.com/mithun/client/kafca"
)

var (
	flag string = "false"
	city string = "Bengaluru"
)

type FuelService interface {
	FuelLog(string, string, bool) error
}

type fuelService struct {
	kafcaClient kafca.Kafca
}

func GetNewFuelService(kafcaClient kafca.Kafca) FuelService {
	return &fuelService{kafcaClient: kafcaClient}
}

// handle service logic and send data to kafca client to publish
func (service *fuelService) FuelLog(newCity string, newFlag string, isvalid bool) error {

	// thsi flag checks if the request is coming from cronjob or api request
	// if true request if from api : update current city and flag accordingly and send data to consumer
	// if false request is from cron : send global value of city and flag
	if isvalid == true {

		// if duplicate value of flag is sent return rerror message indicating to change flag value
		if flag == newFlag {
			return errors.New(fmt.Sprintf("flag already set to %s plz send %s", flag, func() string {
				if flag == "true" {
					return "false"
				} else {
					return "true"
				}
			}()))
		}

		city = newCity
		flag = newFlag

	}

	// call kafca client to publish data
	err := service.kafcaClient.Publish(city + "-" + flag)
	if err != nil {
		return err
	}

	return nil
}
