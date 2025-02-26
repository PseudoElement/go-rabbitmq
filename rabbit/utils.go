package rabbit

import (
	"log"
	"os"
	"strconv"
)

func getRabbitMqUrl() string {
	withDockerString := os.Getenv("WITH_DOCKER")
	withDocker, _ := strconv.ParseBool(withDockerString)
	if withDocker {
		return "amqp://sintol:sintol@rabbitmq:5672/"
	}
	return "amqp://sintol:sintol@localhost:5672/"
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
