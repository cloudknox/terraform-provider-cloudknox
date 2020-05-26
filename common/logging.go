package common

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger
var loggerOnce sync.Once

const (
	output = "info.log"
)

func getLogger() *log.Logger {
	loggerOnce.Do(
		func() {
			/* Initialize Logger */
			logger = log.New()
			// formatter := &log.TextFormatter{
			// 	DisableColors: false,
			// 	FullTimestamp: true,
			// }
			logger.SetFormatter(&log.JSONFormatter{})
			//ogger.SetLevel(log.DebugLevel)

			err := os.Remove(output)

			file, err := os.OpenFile(output, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				logger.Fatal(err)
			}

			logger.SetOutput(file)

			logger.Info("Successfully Created Logger Instance!")

			defer file.Close()
		},
	)
	return logger
}

func GetLogger() *logrus.Logger {
	logger := getLogger()

	file, err := os.OpenFile(output, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Fatal(err)
	}
	logger.SetOutput(file)
	return logger
}
