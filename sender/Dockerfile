FROM golang:latest

WORKDIR /usr/src/app

# Copy the local module directory
COPY rabbit /usr/src/app/rabbit
COPY sender /usr/src/app/sender

# Set working directory inside the container (like $ cd /usr/src/app/sender)
WORKDIR /usr/src/app/sender

# Ensure dependencies resolve correctly
RUN go mod tidy

# Build sender
RUN go build -o sender

EXPOSE 8080

CMD [ "./sender" ]