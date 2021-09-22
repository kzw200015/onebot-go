package onebot

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
)

func DefaultLogger(level logrus.Level) *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetLevel(level)
	logger.SetFormatter(&defaultFormatter{})
	return logger
}

type defaultFormatter struct{}

func (*defaultFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b bytes.Buffer
	t := entry.Time.Format("2006-01-02 15:04:05")
	s := fmt.Sprintf("%s [%s] %s %s\n", t, entry.Level, entry.Caller.Function, entry.Message)
	b.WriteString(s)
	return b.Bytes(), nil
}
