package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(
		&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			PadLevelText:    true,
			ForceColors:     true,
		})
	//logger.SetFormatter(&logrus.JSONFormatter{
	//	TimestampFormat: "2006-01-02 15:04:05",
	//	PrettyPrint:     true,
	//})
}

func GetLogger() *logrus.Logger {
	return logger
}
