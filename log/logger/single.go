package logger

import (
	"errors"
	"gopkg.in/go-mixed/framework.v1/facades/config"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"gopkg.in/go-mixed/framework.v1/log/formatter"
)

type Single struct {
}

func (single *Single) Handle(channel string) (logrus.Hook, error) {
	logPath := config.GetString(channel + ".path")
	if logPath == "" {
		return nil, errors.New("error log path")
	}

	levels := getLevels(config.GetString(channel + ".level"))
	pathMap := lfshook.PathMap{}
	for _, level := range levels {
		pathMap[level] = logPath
	}

	return lfshook.NewHook(
		pathMap,
		&formatter.General{},
	), nil
}

func getLevels(level string) []logrus.Level {
	if level == "panic" {
		return []logrus.Level{
			logrus.PanicLevel,
		}
	}

	if level == "fatal" {
		return []logrus.Level{
			logrus.FatalLevel,
			logrus.PanicLevel,
		}
	}

	if level == "error" {
		return []logrus.Level{
			logrus.ErrorLevel,
			logrus.FatalLevel,
			logrus.PanicLevel,
		}
	}

	if level == "warning" {
		return []logrus.Level{
			logrus.WarnLevel,
			logrus.ErrorLevel,
			logrus.FatalLevel,
			logrus.PanicLevel,
		}
	}

	if level == "info" {
		return []logrus.Level{
			logrus.InfoLevel,
			logrus.WarnLevel,
			logrus.ErrorLevel,
			logrus.FatalLevel,
			logrus.PanicLevel,
		}
	}

	return []logrus.Level{
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}
