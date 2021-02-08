FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go get github.com/githubnemo/CompileDaemon

Run go mod vendor

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main