// Package logger
/**
 * @Author: sunyang
 * @Email: sunyang@hyperchain.cn
 * @Date: 2022/10/12
 */
package logger

import (
	"gitee.com/goweb/tools/viperfile"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var logonce sync.Once

var l *zap.Logger

func Logger() *zap.Logger {
	logonce.Do(func() {
		l = getLogger()
	})
	return l
}

func getLogger() *zap.Logger {
	infoWriter := lumberjack.Logger{
		Filename:   viperfile.DefaultViperString("log.info_file", "log/info.log"), // 日志输出地址
		LocalTime:  viperfile.DefaultViperBool("log.filename_with_time", true),    // 日志文件名时间
		MaxSize:    viperfile.DefaultViperInt("log.file_max_size", 100),           // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: viperfile.DefaultViperInt("log.file_max_backups", 30),         // 日志文件最多保存多少个备份
		MaxAge:     viperfile.DefaultViperInt("log.file_max_age", 30),             // 文件最多保存多少天
		Compress:   viperfile.DefaultViperBool("log.file_compress", true),         // 是否压缩
	}
	errorWriter := lumberjack.Logger{
		Filename:   viperfile.DefaultViperString("log.error_file", "log/error.log"), // 日志输出地址
		LocalTime:  viperfile.DefaultViperBool("log.filename_with_time", true),      // 日志文件名时间
		MaxSize:    viperfile.DefaultViperInt("log.file_max_size", 100),             // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: viperfile.DefaultViperInt("log.file_max_backups", 30),           // 日志文件最多保存多少个备份
		MaxAge:     viperfile.DefaultViperInt("log.file_max_age", 30),               // 文件最多保存多少天
		Compress:   viperfile.DefaultViperBool("log.file_compress", true),           // 是否压缩
	}
	// 日志输出格式
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "",
		MessageKey:     "msg",
		StacktraceKey:  "",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // FullCallerEncoder
		EncodeName:     zapcore.FullNameEncoder,
	}
	consoleConfig := zapcore.EncoderConfig{
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "",
		CallerKey:      "",
		MessageKey:     "msg",
		StacktraceKey:  "",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 根据配置调整日志级别, 支持http接口动态修改zap日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	var l zapcore.Level
	_ = l.UnmarshalText([]byte(viperfile.DefaultViperString("log.console_level", "info")))
	atomicLevel.SetLevel(l)

	// 设置输出源，输出格式，日志等级
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(&infoWriter), zap.InfoLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(&errorWriter), zap.ErrorLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(consoleConfig), zapcore.AddSync(os.Stdout), atomicLevel),
	)
	logger := zap.New(core)

	if viperfile.DefaultViperBool("log.debug", false) {
		// 开启开发模式，堆栈跟踪
		logger.WithOptions(zap.AddCaller(), zap.AddStacktrace(zap.InfoLevel))
	}
	logger.WithOptions(zap.Development())
	return logger
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
