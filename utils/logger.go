package utils

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

type LogService struct {
	logging *logrus.Logger
}

func InitLogger() *LogService {
	logrusLogger := logrus.New()
	logrusLogger.SetOutput(os.Stdout)
	logrusLogger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2000-00-00 00:00:00.000"})
	//logrusLogger.SetFormatter(&logrus.TextFormatter{TimestampFormat: ""})
	level, _ := logrus.ParseLevel(Cfg.LogLevel)
	logrusLogger.SetLevel(level)
	return &LogService{
		logging: logrusLogger,
	}
}

func (l LogService) Log() logrus.FieldLogger {
	return l.logging
}

func (l LogService) LogWithContext(ctx context.Context) logrus.FieldLogger {
	return l.logging.WithFields(logrus.Fields{
		"CORRELATION_ID": ctx.Value("correlationId").(string),
		// Any other items to trace each request
	})
}
