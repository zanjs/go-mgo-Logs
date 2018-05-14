package tools

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

//Log is
var Log *logrus.Logger

// NewLogger is
func NewLogger() *logrus.Logger {
	if Log != nil {
		return Log
	}
	// path := "./go.log"
	logf, _ := rotatelogs.New(
		"./log/%Y-%m-%d.go.log",
		rotatelogs.WithLinkName("./access_log"),
		rotatelogs.WithMaxAge(168*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)

	Log = logrus.New()

	Log.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  logf,
			logrus.ErrorLevel: logf,
		},
		&logrus.JSONFormatter{},
	))

	return Log
}

func init() {
	Log = NewLogger()
}
