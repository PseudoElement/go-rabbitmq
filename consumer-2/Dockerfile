FROM golang:latest

WORKDIR /usr/src/app

# Copy the local module directory
COPY rabbit /usr/src/app/rabbit
COPY consumer-2 /usr/src/app/consumer-2

# Set working directory inside the container (like $ cd /usr/src/app/sender)
WORKDIR /usr/src/app/consumer-2

# Ensure dependencies resolve correctly
RUN go mod tidy

# Build consumer
RUN go build -o consumer-2

EXPOSE 8082

CMD [ "./consumer-2" ]