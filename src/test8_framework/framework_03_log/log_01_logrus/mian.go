package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)


func init() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(log.DebugLevel)

	log.Info("This is an info message")
	log.WithFields(log.Fields{
		"username": "tom",
		"age":      30,
	}).Warn("This is a warning with fields")
}
