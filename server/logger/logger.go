package logger

import (
	"github.com/Sirupsen/logrus"
	"os"
)

const (
	logsFileName = "./var/log/run.log"
)

type Logger interface {
	Info(message string, fields Fields)
	Error(message string, fields Fields)
}

type Fields logrus.Fields

type logger struct {
	prefix string
}

func New(prefix string) Logger {
	file, err := os.OpenFile(logsFileName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	logrus.SetOutput(file)

	return &logger{
		prefix: prefix,
	}
}

func (t *logger) Info(message string, fields Fields) {
	if t.prefix != "" {
		fields["prefix"] = t.prefix
	}
	logrus.WithFields(logrus.Fields(fields)).Info(message)
}

func (t *logger) Error(message string, fields Fields) {
	if t.prefix != "" {
		fields["prefix"] = t.prefix
	}
	logrus.WithFields(logrus.Fields(fields)).Error(message)
}
