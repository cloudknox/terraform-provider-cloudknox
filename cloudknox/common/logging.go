package common

import (
	"fmt"
	"os"
	"sync"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

var customLogger CustomLogger
var loggerOnce sync.Once

const (
	output = "info.log"
)

func getLogger() CustomLogger {
	loggerOnce.Do(
		func() {
			/* Initialize Logger */

			err := os.Remove(output)

			file, err := os.OpenFile(output, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err == nil {
				customLogger.logger = log.NewLogfmtLogger(file)
				customLogger.logger = level.NewFilter(customLogger.logger, level.AllowAll())
				customLogger.logger = log.With(customLogger.logger, "time", log.DefaultTimestampUTC)
				customLogger.Info("msg", "Successfully Created Logger Instance!")
			} else {
				fmt.Println("Unable to begin logging")
			}

		},
	)
	return customLogger
}

func GetLogger() CustomLogger {
	logger := getLogger()

	// file, err := os.OpenFile(output, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	logger.Fatal(err)
	// }
	// logger.SetOutput(file)
	return logger
}

type CustomLogger struct {
	logger log.Logger
}

func (clog CustomLogger) Info(args ...interface{}) {
	level.Info(clog.logger).Log(args...)
}

func (clog CustomLogger) Debug(args ...interface{}) {
	level.Debug(clog.logger).Log(args...)
}

func (clog CustomLogger) Warn(args ...interface{}) {
	level.Warn(clog.logger).Log(args...)
}

func (clog CustomLogger) Error(args ...interface{}) {
	level.Error(clog.logger).Log(args)
}
