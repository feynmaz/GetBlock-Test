package logger

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	ErrLogLevelUnknown = errors.New("log level unknown")
)

func InitLogrus(logLevel string, logJson bool) error {
	if logJson {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	logrus.SetOutput(os.Stdout)

	var level logrus.Level
	switch strings.ToLower(logLevel) {
	case "debug":
		level = logrus.DebugLevel

	case "info":
		level = logrus.InfoLevel

	case "warn":
		level = logrus.WarnLevel

	default:
		return fmt.Errorf("%w: %s", ErrLogLevelUnknown, logLevel)
	}
	logrus.SetLevel(level)
	return nil
}
