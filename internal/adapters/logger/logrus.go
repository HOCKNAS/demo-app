package logger

import (
	"os"

	"github.com/HOCKNAS/demo-app/internal/core/ports"
	"github.com/HOCKNAS/demo-app/pkg/log/logrus_log"
	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger(config *logrus.TextFormatter) ports.Logger {

	logger := logrus_log.NewLogger()
	logger.Out = os.Stdout
	logger.SetReportCaller(true)
	logger.SetFormatter(config)

	return &logrusLogger{
		logger: logger,
	}
}

func convertArgsToFields(args ...interface{}) logrus.Fields {
	fields := logrus.Fields{}
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			key, ok := args[i].(string)
			if !ok {
				fields["loggingError"] = "La clave de argumento no es una cadena"
				continue
			}
			fields[key] = args[i+1]
		} else {
			fields["unpairedArg"] = args[i]
		}
	}
	return fields
}

func (l *logrusLogger) Debug(msg string, args ...interface{}) {
	l.logger.WithFields(convertArgsToFields(args...)).Debug(msg)
}

func (l *logrusLogger) Info(msg string, args ...interface{}) {
	l.logger.WithFields(convertArgsToFields(args...)).Info(msg)
}

func (l *logrusLogger) Warn(msg string, args ...interface{}) {
	l.logger.WithFields(convertArgsToFields(args...)).Warn(msg)
}

func (l *logrusLogger) Error(msg string, args ...interface{}) {
	l.logger.WithFields(convertArgsToFields(args...)).Error(msg)
}
