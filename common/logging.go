package common

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var loggerOnce sync.Once

const (
	output = "info.log"
)

func getLogger() *logrus.Logger {
	loggerOnce.Do(
		func() {
			/* Initialize Logger */
			logger = logrus.New()
			formatter := &logrus.TextFormatter{
				DisableColors:   true,
				TimestampFormat: "2006-01-02 15:04:05",
				FullTimestamp:   true,
			}
			logger.SetFormatter(formatter)

			err := os.Remove(output)

			file, err := os.OpenFile(output, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				logger.Fatal(err)
			}

			logger.SetOutput(file)

			logger.Info("Successfully Created Logger Instance!")

		},
	)
	return logger
}

func GetLogger() *logrus.Logger {
	logger := getLogger()

	// file, err := os.OpenFile(output, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	logger.Fatal(err)
	// }
	// logger.SetOutput(file)
	return logger
}
