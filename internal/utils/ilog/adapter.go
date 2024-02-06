package ilog

import (
	"context"
	"io"
	"time"

	"github.com/sirupsen/logrus"
)

type LogrusWrapper struct {
	base *logrus.Entry
}

func (l *LogrusWrapper) Log(i ...interface{}) {
	l.base.Println(i...)
}

func (l *LogrusWrapper) WithField(key string, value interface{}) Logger {
	return fromEntry(l.base.WithField(key, value))
}

func (l *LogrusWrapper) WithFields(fields logrus.Fields) Logger {
	return fromEntry(l.base.WithFields(fields))
}

func (l *LogrusWrapper) WithError(err error) Logger {
	return fromEntry(l.base.WithError(err))
}

func (l *LogrusWrapper) WithContext(ctx context.Context) Logger {
	return fromEntry(l.base.WithContext(ctx))
}

func (l *LogrusWrapper) WithTime(t time.Time) Logger {
	return fromEntry(l.base.WithTime(t))
}

func (l *LogrusWrapper) Debugf(format string, args ...interface{}) {
	l.base.Debugf(format, args...)
}

func (l *LogrusWrapper) Infof(format string, args ...interface{}) {
	l.base.Infof(format, args...)
}

func (l *LogrusWrapper) Printf(format string, args ...interface{}) {
	l.base.Printf(format, args...)
}

func (l *LogrusWrapper) Warningf(format string, args ...interface{}) {
	l.base.Warningf(format, args...)
}

func (l *LogrusWrapper) Errorf(format string, args ...interface{}) {
	l.base.Errorf(format, args...)
}

func (l *LogrusWrapper) Fatalf(format string, args ...interface{}) {
	l.base.Fatalf(format, args...)
}

func (l *LogrusWrapper) Panicf(format string, args ...interface{}) {
	l.base.Panicf(format, args...)
}

func (l *LogrusWrapper) Debug(args ...interface{}) {
	l.base.Debug(args...)
}

func (l *LogrusWrapper) Info(args ...interface{}) {
	l.base.Info(args...)
}

func (l *LogrusWrapper) Print(args ...interface{}) {
	l.base.Print(args...)
}

func (l *LogrusWrapper) Warning(args ...interface{}) {
	l.base.Warning(args...)
}

func (l *LogrusWrapper) Error(args ...interface{}) {
	l.base.Error(args...)
}

func (l *LogrusWrapper) Fatal(args ...interface{}) {
	l.base.Fatal(args...)
}

func (l *LogrusWrapper) Panic(args ...interface{}) {
	l.base.Panic(args...)
}

func (l *LogrusWrapper) WriterLevel(level logrus.Level) io.Writer {
	return l.base.WriterLevel(level)
}

func NewLogrusWrapper(logger *logrus.Logger) Logger {
	return &LogrusWrapper{logrus.NewEntry(logger)}
}

func fromEntry(entry *logrus.Entry) Logger {
	return &LogrusWrapper{entry}
}
