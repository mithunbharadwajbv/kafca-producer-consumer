FROM golang:latest

WORKDIR /app

COPY ./ /app

Ad

RUN go get github.com/githubnemo/CompileDaemon

Run go mod vendor

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main