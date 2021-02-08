package main

import (
	"github.com/mithun/app"
	"github.com/mithun/utils"
)

func main() {

	//cron job to push data every 2 minutes
	go utils.Schedulecal()

	app.StartApplication()

}
