package log

import "github.com/sirupsen/logrus"

// Logger is used here to de-couple logging framework
type Logger interface {
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})

	// Fatal must execute os.Exit(1) after logging
	Fatalf(template string, args ...interface{})

	// Panic must execute panic() after logging
	Panicf(template string, args ...interface{})
}

// NewLogger is used to create a new logger
func NewLogger() Logger {
	return logrus.New()
}
