FROM golang:1.17.2-alpine

WORKDIR /app
COPY . /app

RUN go get -d -v ./...

RUN go build main.go

CMD ./main