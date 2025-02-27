FROM golang:latest

WORKDIR /usr/src/app

COPY sender/go.mod sender/go.sum ./

# Copy the local module directory (assuming it's in your project structure)
COPY rabbit /usr/src/app/rabbit

RUN go mod download && go mod verify && go mod tidy

COPY sender/*.go .

RUN go build -o sender

EXPOSE 8080

CMD [ "./sender" ]