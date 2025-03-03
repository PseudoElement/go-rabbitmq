FROM golang:latest

WORKDIR /usr/src/app

# Copy the local module directory
COPY rabbit /usr/src/app/rabbit
COPY consumer /usr/src/app/consumer

# Set working directory inside the container (like $ cd /usr/src/app/sender)
WORKDIR /usr/src/app/consumer

# Ensure dependencies resolve correctly
RUN go mod tidy

# Build consumer
RUN go build -o consumer

EXPOSE 8081

CMD [ "./consumer" ]