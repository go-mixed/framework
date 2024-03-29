package logger

import (
	"errors"
	rotatelogs "github.com/goravel/file-rotatelogs/v2"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-mixed/framework.v1/facades/config"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"gopkg.in/go-mixed/framework.v1/log/formatter"
)

type Daily struct {
}

func (daily *Daily) Handle(channel string) (logrus.Hook, error) {
	var hook logrus.Hook
	logPath := config.GetString(channel + ".path")
	if logPath == "" {
		return hook, errors.New("error log path")
	}

	ext := path.Ext(logPath)
	logPath = strings.ReplaceAll(logPath, ext, "")

	writer, err := rotatelogs.New(
		logPath+"-%Y-%m-%d"+ext,
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		rotatelogs.WithRotationCount(uint(config.GetInt(channel+".days"))),
	)
	if err != nil {
		return hook, errors.New("Config local file system for logger error: " + err.Error())
	}
	writer2 := os.Stdout
	levels := getLevels(config.GetString(channel + ".level"))
	writerMap := lfshook.WriterMap{}
	for _, level := range levels {
		writerMap[level] = io.MultiWriter(writer, writer2)
	}

	return lfshook.NewHook(
		writerMap,
		&formatter.General{},
	), nil
}
