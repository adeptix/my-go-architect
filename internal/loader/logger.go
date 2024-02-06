package loader

import (
	"os"

	"github.com/sirupsen/logrus"

	"mycleanarch/internal/utils/ilog"
)

func InitLogger() ilog.Logger {
	level := os.Getenv("LOGGER_LEVEL")
	if level == "" {
		level = "debug"
	}

	logger := logrus.New()

	logLevel, err := logrus.ParseLevel(level)
	if err == nil {
		logger.SetLevel(logLevel)
	}

	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	//InitLoggerSentryHook(logger)
	//
	//klog.ClearLogger()
	//klog.SetOutputBySeverity("INFO", logger.WriterLevel(logrus.InfoLevel))
	//klog.SetOutputBySeverity("WARNING", logger.WriterLevel(logrus.WarnLevel))
	//klog.SetOutputBySeverity("ERROR", logger.WriterLevel(logrus.ErrorLevel))
	//klog.SetOutputBySeverity("FATAL", logger.WriterLevel(logrus.FatalLevel))

	return ilog.NewLogrusWrapper(logger)
}

//func InitLoggerSentryHook(l *logrus.Logger) {
//	client := sentry.CurrentHub().Client()
//	if client == nil {
//		return
//	}
//
//	levels := []logrus.Level{
//		logrus.FatalLevel,
//		logrus.PanicLevel,
//		logrus.ErrorLevel,
//		logrus.WarnLevel,
//		logrus.InfoLevel,
//		// logrus.DebugLevel,
//	}
//
//	hook := sentrylogrus.NewFromClient(levels, client)
//	hook.SetContextHub(sentrylogrus.DefaultContextExtractor)
//
//	l.Hooks.Add(hook)
//}
