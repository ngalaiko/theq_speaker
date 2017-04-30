package logger

import "github.com/Sirupsen/logrus"

type Logger interface {
	Info(message string, fields Fields)
	Error(message string, fields Fields)
}

type Fields logrus.Fields

type logger struct {
	prefix string
}

func New(prefix string) Logger {
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
