package log

import (
	"context"
	"github.com/gookit/color"
	"github.com/sirupsen/logrus"

	"gopkg.in/go-mixed/framework.v1/contracts/log"
)

type Logger struct {
	instance *logrus.Logger
	log.Writer
}

func WrapLogger(writer log.Writer) *Logger {
	return &Logger{
		Writer: writer,
	}
}

func NewLogger(ctx context.Context, channelName string) (*Logger, error) {
	instance, err := newLogrus(channelName)
	if err != nil {
		color.Redln("Initialize log error: " + err.Error())
		return nil, err
	}

	return &Logger{
		instance: instance,
		Writer:   NewWriter(instance.WithContext(ctx)),
	}, nil
}

func (r *Logger) WithContext(ctx context.Context) log.Writer {
	switch r.Writer.(type) {
	case *Writer:
		return NewWriter(r.instance.WithContext(ctx))
	default:
		return r.Writer
	}
}

func newLogrus(channelName string) (*logrus.Logger, error) {
	instance := logrus.New()
	instance.SetLevel(logrus.DebugLevel)

	return instance, registerHook(instance, channelName)
}
