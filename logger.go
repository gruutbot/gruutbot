package gruutbot

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Trace(v ...interface{})
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
	Panic(v ...interface{})
	Tracef(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
	Panicf(format string, v ...interface{})
	Traceln(v ...interface{})
	Debugln(v ...interface{})
	Infoln(v ...interface{})
	Warnln(v ...interface{})
	Errorln(v ...interface{})
	Fatalln(v ...interface{})
	Panicln(v ...interface{})
}

type LogLevel string

const (
	TraceLevel  LogLevel = "trace"
	DebugLevel  LogLevel = "debug"
	InfoLevel   LogLevel = "info"
	WarnLevel   LogLevel = "warn"
	ErrorLevel  LogLevel = "error"
	FatalLevel  LogLevel = "fatal"
	logLevelKey          = "log_level"
)

func logrusLogger(level LogLevel) (lf *logrus.Entry) {
	l := logrus.New()

	if gviper.IsSet(logLevelKey) {
		level = LogLevel(gviper.GetString(logLevelKey))
	}

	l.SetLevel(logrusLevel(level))
	lf = l.WithField("lib", "gruutbot")
	lf.Infoln("Log level set to", l.Level)

	return
}

func logrusLevel(level LogLevel) logrus.Level {
	level = levelToLower(level)

	switch level {
	case TraceLevel:
		return logrus.TraceLevel
	case DebugLevel:
		return logrus.DebugLevel
	case InfoLevel:
		return logrus.InfoLevel
	case WarnLevel:
		return logrus.WarnLevel
	case ErrorLevel:
		return logrus.ErrorLevel
	case FatalLevel:
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}

func levelToLower(l LogLevel) LogLevel {
	s := string(l)
	s = strings.ToLower(s)

	return LogLevel(s)
}
