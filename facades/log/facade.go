package log

import (
	"context"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/log"
)

func getLogger() log.ILog {
	return container.MustMakeAs("log", log.ILog(nil))
}

func WithContext(ctx context.Context) log.Writer {
	return getLogger().WithContext(ctx)
}

func Channel(name string) log.ILog {
	return getLogger().Channel(name)
}

func Debug(args ...any) {
	getLogger().Debug(args...)
}

func Debugf(format string, args ...any) {
	getLogger().Debugf(format, args...)
}

func Info(args ...any) {
	getLogger().Info(args...)
}

func Infof(format string, args ...any) {
	getLogger().Infof(format, args...)
}

func Warning(args ...any) {
	getLogger().Warning(args...)
}

func Warningf(format string, args ...any) {
	getLogger().Warningf(format, args...)
}

func Error(args ...any) {
	getLogger().Error(args...)
}

func Errorf(format string, args ...any) {
	getLogger().Errorf(format, args...)
}

func Fatal(args ...any) {
	getLogger().Fatal(args...)
}

func Fatalf(format string, args ...any) {
	getLogger().Fatalf(format, args...)
}

func Panic(args ...any) {
	getLogger().Panic(args...)
}

func Panicf(format string, args ...any) {
	getLogger().Panicf(format, args...)
}
