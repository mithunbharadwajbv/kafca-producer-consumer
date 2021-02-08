package fulelog

import (
	"consumer/client/redis"
	"consumer/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	timeFlag bool = false
	start         = time.Now()

	logger     = utils.GetLogger()
	filelogger = utils.GetFileLogger()
)

func Log(input string) {

	logger.Info(fmt.Sprintf("received data : %s \n", input))

	var redisClient = redis.GetNewRedisClient()

	temp := strings.Split(input, "-")
	city := temp[0]
	flag, _ := strconv.ParseBool(temp[1])

	//Get petrol price value from redis(chache) for a perticular city
	petrolCost := func() float64 {
		var costString string
		var err error
		for i := 0; i < 2; i++ {
			costString, err = redisClient.GetValue(city)
			if err != nil {
				utils.SaveCityPrice() //call api to populate redis
			} else {
				break
			}
		}
		if err != nil {
			panic(fmt.Sprintf("No price visibility for City %s in cache or on api", city))
		}
		cost, _ := strconv.ParseFloat(costString, 64)
		return cost
	}()

	if flag == true {
		if timeFlag == false {

			timeFlag = true
			start = time.Now()

		} else {
			return
		}
	}

	if flag == false {
		if timeFlag == true {

			timeFlag = false
			elapsed := time.Since(start)

			if elapsed.Seconds() > 2.0 {
				filelogger.Info(fmt.Sprintf("time elapsed : %f , cost in %s : %f , totalCost : %v \n", elapsed.Seconds(), city, petrolCost, elapsed.Seconds()*petrolCost/30.0))
			}
		} else {
			return
		}
	}

}
