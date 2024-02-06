package ilog

import (
	"context"
	"io"
	"time"

	"github.com/sirupsen/logrus"
)

type S3Logger interface {
	Log(i ...interface{})
}

type Logger interface {
	WithField(key string, value interface{}) Logger
	WithFields(fields logrus.Fields) Logger
	WithError(err error) Logger
	WithContext(ctx context.Context) Logger
	WithTime(t time.Time) Logger

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	WriterLevel(level logrus.Level) io.Writer

	S3Logger
}
