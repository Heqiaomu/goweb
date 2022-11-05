// Package logger
/**
 * @Author: sunyang
 * @Email: sunyang@hyperchain.cn
 * @Date: 2022/10/12
 */
package logger

import (
	"fmt"
	"gitee.com/goweb/config"
	"gitee.com/goweb/tools/file"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var logonce sync.Once

var l *zap.Logger

func Logger() *zap.Logger {
	if l != nil {
		return l
	}
	logonce.Do(func() {
		l = getLogger()
	})
	return l
}

func getLogger() *zap.Logger {
	if ok, _ := file.PathExists(config.GetConfig().Zap.Director); !ok {
		fmt.Printf("create %v directory\n", config.GetConfig().Zap.Director)
		_ = os.Mkdir(config.GetConfig().Zap.Director, os.ModePerm)
	}
	cores := GetZapCores()

	l = zap.New(zapcore.NewTee(cores...))

	if config.GetConfig().Zap.ShowLine {
		l = l.WithOptions(zap.AddCaller())
	}
	return l
}

type Level string

const (
	DebugLevel Level = "debug"
	InfoLevel  Level = "info"
	WarnLevel  Level = "warn"
	ErrorLevel Level = "error"
	PanicLevel Level = "panic"
	FatalLevel Level = "fatal"
)

func Log(format string, fields ...zap.Field) {
	l.Info(format, fields...)
}
func Debug(format string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Debug(format, fields...)
	// TODO 发往日志系统
}
func Debugf(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Debugf(template, args...)
	// TODO 发往日志系统
}
func Debugw(template string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Debug(template, fields...)
	// TODO 发往日志系统
}

func Info(format string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Info(format, fields...)
	// TODO 发往日志系统
}
func Infof(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Infof(template, args...)
	// TODO 发往日志系统
}
func Infow(template string, keysAndValues ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Infow(template, keysAndValues...)
	// TODO 发往日志系统
}

func Warn(format string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Warn(format, fields...)
	// TODO 发往日志系统
}
func Warnf(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Warnf(template, args...)
	// TODO 发往日志系统
}
func Warnw(template string, keysAndValues ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Warnw(template, keysAndValues...)
	// TODO 发往日志系统
}

func Error(format string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Error(format, fields...)
	// TODO 发往日志系统
}
func Errorf(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Errorf(template, args...)
	// TODO 发往日志系统
}
func Errorw(template string, keysAndValues ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Errorw(template, keysAndValues...)
	// TODO 发往日志系统
}

func Panic(format string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Panic(format, fields...)
	// TODO 发往日志系统
}
func Panicf(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Panicf(template, args...)
	// TODO 发往日志系统
}
func Panicw(template string, keysAndValues ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Panicw(template, keysAndValues...)
	// TODO 发往日志系统
}

func Falal(format string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Fatal(format, fields...)
	// TODO 发往日志系统
}
func Fatalf(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Fatalf(template, args...)
	// TODO 发往日志系统
}
func Fatalw(template string, keysAndValues ...interface{}) {
	funcName, caller := stackTrace()
	l.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Fatalw(template, keysAndValues...)
	// TODO 发往日志系统
}

func stackTrace() (funcName, caller string) {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, false)]
	stacks := strings.Split(string(buf), "\n")

	funcName = stacks[5]
	funcName = funcName[strings.LastIndex(funcName, "/")+1 : strings.LastIndex(funcName, "(")]
	callers := strings.Split(stacks[6][1:], " ")

	return funcName, callers[0]
}

func constructFieldMap(level, funcName, caller string, fields []zap.Field) map[string]interface{} {
	m := make(map[string]interface{})
	m["level"] = level
	m["funcName"] = funcName
	m["caller"] = caller
	for _, f := range fields {
		m[f.Key] = f.String
	}
	return m
}

func constructArrayMap(level, funcName, caller string, fields []interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	m["level"] = level
	m["funcName"] = funcName
	m["caller"] = caller
	for i, f := range fields {
		m[strconv.Itoa(i)] = f
	}
	return m
}

// Sugar return zap SugaredLogger instance
func Sugar() *zap.SugaredLogger {
	return l.Sugar()
}

// Named return zap instance
func Named(s string) *zap.Logger {
	return l.Named(s)
}

// WithOptions log with option
func WithOptions(opts ...zap.Option) *zap.Logger {
	return l.WithOptions(opts...)
}

// With log with field
func With(fields ...zap.Field) *zap.Logger {
	return l.With(fields...)
}

// Check level check
func Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	return l.Check(lvl, msg)
}

// Sync log sync
func Sync() error {
	return l.Sync()
}

// Core return log core
func Core() zapcore.Core {
	return l.Core()
}
