package system

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func LogInit() {
	logFile, _ := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logDest := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(logDest)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006/01/02 15:04:05"})
}
