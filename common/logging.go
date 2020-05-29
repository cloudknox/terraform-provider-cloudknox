package common

import (
	"fmt"
	"os"
	"sync"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

var logger log.Logger
var loggerOnce sync.Once

const (
	output = "info.log"
)

func getLogger() log.Logger {
	loggerOnce.Do(
		func() {
			/* Initialize Logger */

			err := os.Remove(output)

			file, err := os.OpenFile(output, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err == nil {
				logger = log.NewLogfmtLogger(file)
				logger = level.NewFilter(logger, level.AllowInfo())
				logger = log.With(logger, "ts", log.DefaultTimestampUTC)
				level.Info(logger).Log("msg", "Successfully Created Logger Instance!")
			} else {
				fmt.Println("Unable to begin logging")
			}

		},
	)
	return logger
}

func GetLogger() log.Logger {
	logger := getLogger()

	// file, err := os.OpenFile(output, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	logger.Fatal(err)
	// }
	// logger.SetOutput(file)
	return logger
}
