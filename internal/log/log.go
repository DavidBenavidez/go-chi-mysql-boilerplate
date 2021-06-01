package log

import "github.com/sirupsen/logrus"

var log *logrus.Entry = logrus.NewEntry(logrus.StandardLogger())

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Infoln(args ...interface{}) {
	log.Infoln(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Fatalln(format string) {
	log.Fatalln(format)
}
