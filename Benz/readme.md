# <TODO>

1. Adding Jwt Token for Api Authontication
2. dockerising the application

# Task Accomplished

1. Project contains 2 Microservices (producer and consumer) communicating on a event driven moden using kafca(github.com/segmentio/kafka-go)

2. Cache implementation has been done to save cost of petron in various cities of karnataka using redis (every 24 hrs cache data is gonna be earased and new new data will be populated)
redis package used : "github.com/go-redis/redis/v8"
3rd part api used :https://daily-fuel-prices-india.p.rapidapi.com/api/proxy/hp/states/KA
note : if random city is encountered our consimer is gonna panic(design desission)

3. Cron job is implemented on the producer side to push latest status of the location and fuel fill status(true or false), Scheduling has been done for 2 minutes

4. Code is highly modularised with hexagonal architecture

5. configration parametes are coganised and accessed through viper


# Producer Structure

├── app
│   └── app.go
├── client
│   └── kafca
│       ├── config.go
│       └── fuelLogProducer.go
├── config
│   ├── config.go
│   └── conf.json
├── controller
│   └── fuelcontroller.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── models
│   └── models.go
├── service
│   └── fuelService.go
└── utils
    ├── cron.go
    └── logger.go


1. gin has been used to create apis

2. code structure is maintained such that any client can be replaced without affecting other code component by much

3. uber/zap has been used for high efficient logging

[run-pplication]

Run In terminal-1

1. cd to directory which contains producers main.go
2. docker-compose up -d
2. go run main.go

# Consumer Structure

├── app
│   └── fulelog
│       ├── consumer.go
│       └── log.go
├── client
│   └── redis
│       ├── config.go
│       └── redis.go
├── config
│   ├── config.go
│   └── conf.json
├── go.mod
├── go.sum
├── main.go
└── utils
    ├── logger.go
    └── util.go


1. Core logic is written inside app/fuellog which listens to all the messages that producer has sent and take decissions accordingly

2. utils/util.go contains 3rd party api to popullate redis cache when data is not found in cache(this automatically happens through function call no cron job has bee kept to accomplish this functionality)

[run-application]

Run In terminal-1

1. cd to directory which contains consumers main.go
2. go run main.go


# Run on 3rd terminal

1. Post request to send notification of fuel fill has been initiated

curl --location --request POST 'localhost:8080/logfuel' \
--header 'Content-Type: application/json' \
--data-raw '{
    "city" : "Bengaluru",
    "flag" : "true",
    "mobile" : true
}'

2. wait for sometime (more that 10 second)

3. Post request to send notification of fuel fill has been ended

curl --location --request POST 'localhost:8080/logfuel' \
--header 'Content-Type: application/json' \
--data-raw '{
    "city" : "Bengaluru",
    "flag" : "false",
    "mobile" : true
}'

4. repeat above 3 steps multiple times to get multiple entries in logfile


[note]

1. all the logs will be created in the main folder naming fuelog.json











