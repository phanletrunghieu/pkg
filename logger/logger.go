package logger

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	field  logrus.FieldLogger
	level  logrus.Level
	fields logrus.Fields
}

type Option func(*Logger)

var errMissingValue = errors.New("(MISSING)")

func NewLogger(fullTimestamp bool) *Logger {
	l := logrus.New()

	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: fullTimestamp,
	})

	return &Logger{
		field:  l,
		level:  logrus.InfoLevel,
		fields: logrus.Fields{},
	}
}

func WithLevel(level logrus.Level) Option {
	return func(c *Logger) {
		c.level = level
	}
}

func WithFields(keyvals ...interface{}) Option {
	return func(c *Logger) {
		for i := 0; i < len(keyvals); i += 2 {
			if i+1 < len(keyvals) {
				c.fields[fmt.Sprint(keyvals[i])] = keyvals[i+1]
			} else {
				c.fields[fmt.Sprint(keyvals[i])] = errMissingValue
			}
		}
	}
}

func (l *Logger) Option(options ...Option) *Logger {
	for _, optFunc := range options {
		optFunc(l)
	}

	return l
}

func (l *Logger) Entry(options ...Option) *logrus.Entry {
	return l.field.WithFields(l.fields)
}

func (l *Logger) Log(args ...interface{}) {
	switch l.level {
	case logrus.InfoLevel:
		l.field.WithFields(l.fields).Info(args)
	case logrus.ErrorLevel:
		l.field.WithFields(l.fields).Error(args)
	case logrus.DebugLevel:
		l.field.WithFields(l.fields).Debug(args)
	case logrus.WarnLevel:
		l.field.WithFields(l.fields).Warn(args)
	case logrus.TraceLevel:
		l.field.WithFields(l.fields).Trace(args)
	default:
		l.field.WithFields(l.fields).Print(args)
	}
}
