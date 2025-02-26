FROM golang:latest

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify && go mod tidy

COPY . .

RUN go build -o go-rabbitmq

EXPOSE 8080

CMD [ "./go-rabbitmq" ]