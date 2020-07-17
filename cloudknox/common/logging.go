package common

import (
	"fmt"
	"os"
	"sync"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

/* Private Variables */
var customLogger CustomLogger
var loggerOnce sync.Once

const (
	output = "/var/log/cloudknox/application.log"
)

/* Private Functions */
func getLogger() CustomLogger {
	loggerOnce.Do(
		func() {
			/* Initialize Logger */

			err := os.Remove(output)

			file, err := os.OpenFile(output, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err == nil {
				customLogger.logger = log.NewLogfmtLogger(file)
				customLogger.logger = level.NewFilter(customLogger.logger, level.AllowInfo())
				customLogger.logger = log.With(customLogger.logger, "time", log.DefaultTimestampUTC)
				customLogger.Info("msg", "successfully created logger instance!")
			} else {
				fmt.Println("Unable to begin logging", err)
			}

		},
	)
	return customLogger
}

/* Public Functions */

// GetLogger returns a logger wrapper that incorporates level logging
func GetLogger() CustomLogger {
	return getLogger()
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
	level.Error(clog.logger).Log(args...)
}

func (clog CustomLogger) Fatal(args ...interface{}) {
	level.Error(clog.logger).Log(args...)
	os.Exit(1)
}
