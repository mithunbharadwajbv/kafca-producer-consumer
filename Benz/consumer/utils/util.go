package utils

import (
	"consumer/client/redis"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Fuel struct {
	TownName string `json:"townname"`
	Petrol   string `json:"petrol"`
}

func SaveCityPrice() {

	var redisClient = redis.GetNewRedisClient()

	var fuels []Fuel
	url := "https://daily-fuel-prices-india.p.rapidapi.com/api/proxy/hp/states/KA"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "867002c364msh09e72727dd3bbbfp13209ejsnc16d85737dfa")
	req.Header.Add("x-rapidapi-host", "daily-fuel-prices-india.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(body, &fuels)
	if err != nil {
		fmt.Printf(err.Error())

	}

	fmt.Println(fuels)

	for _, v := range fuels {
		redisClient.SetDataWithExpiry(v.TownName, v.Petrol)
	}

}
