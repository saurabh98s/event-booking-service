package logger

import (
	"github.com/sirupsen/logrus"
)

// Log is the exported object for logging
var Log *logrus.Logger

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	Log = logrus.New()
}
