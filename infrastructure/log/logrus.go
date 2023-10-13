package log

import (
	"fmt"
	"nu/corpus-reader/adapter/logger"
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger(logPrefix string) logger.Logger {
	log := logrus.New()
  log.SetReportCaller(true)
	log.SetFormatter(&logrus.TextFormatter{
    FullTimestamp: true,
    PadLevelText: true,
     CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {                                                     
        fileName := fmt.Sprintf(" [%v] ", logPrefix) + path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)       
        return "", fileName                                                      
     },     
  })

	return &logrusLogger{logger: log}
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *logrusLogger) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args...)
}

func (l *logrusLogger) WithFields(fields logger.Fields) logger.Logger {
	return &logrusLogEntry{
		entry: l.logger.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *logrusLogger) WithError(err error) logger.Logger {
	return &logrusLogEntry{
		entry: l.logger.WithError(err),
	}
}

type logrusLogEntry struct {
	entry *logrus.Entry
}

func (l *logrusLogEntry) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l *logrusLogEntry) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

func (l *logrusLogEntry) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

func (l *logrusLogEntry) Fatalln(args ...interface{}) {
	l.entry.Fatalln(args...)
}

func (l *logrusLogEntry) WithFields(fields logger.Fields) logger.Logger {
	return &logrusLogEntry{
		entry: l.entry.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *logrusLogEntry) WithError(err error) logger.Logger {
	return &logrusLogEntry{
		entry: l.entry.WithError(err),
	}
}

func convertToLogrusFields(fields logger.Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, field := range fields {
		logrusFields[index] = field
	}

	return logrusFields
}
