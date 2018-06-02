package log

import "github.com/sirupsen/logrus"

var logDefault = logrus.New()

type Fields map[string]interface{}

func Infoln(args ...interface{}) {
	logDefault.Infoln(args...)
}
func Infof(format string, args ...interface{}) {
	logDefault.Infof(format, args...)
}

func Debugln(args ...interface{}) {
	logDefault.Debugln(args...)
}
func Debugf(format string, args ...interface{}) {
	logDefault.Debugf(format, args...)
}

func Fatalln(args ...interface{}) {
	logDefault.Fatalln(args...)
}
func Fatalf(format string, args ...interface{}) {
	logDefault.Fatalf(format, args...)
}

func Errorln(args ...interface{}) {
	logDefault.Errorln(args...)
}

func Errorf(format string, args ...interface{}) {
	logDefault.Errorf(format, args...)
}

func WithFields(f Fields) *logrus.Entry {
	return logDefault.WithFields(logrus.Fields(f))
}
