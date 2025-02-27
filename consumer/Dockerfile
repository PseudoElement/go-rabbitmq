FROM golang:latest

WORKDIR /usr/src/app

COPY consumer/go.mod consumer/go.sum ./

# Copy the local module directory (assuming it's in your project structure)
COPY rabbit /usr/src/app/rabbit

RUN go mod download && go mod verify && go mod tidy

COPY consumer/*.go .

RUN go build -o consumer

EXPOSE 8081

CMD [ "./consumer" ]