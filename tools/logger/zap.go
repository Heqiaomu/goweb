package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sun-iot/goweb/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

// GetEncoder 获取 zapcore.Encoder
func GetEncoder() zapcore.Encoder {
	if config.GetConfig().Zap.Format == "json" {
		return zapcore.NewJSONEncoder(GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(GetEncoderConfig())
}

// GetEncoderConfig 获取zapcore.EncoderConfig
func GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "caller",
		StacktraceKey: config.GetConfig().Zap.StacktraceKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   config.GetConfig().Zap.ZapEncodeLevel(),
		//EncodeTime:     CustomTimeEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

// GetEncoderCore 获取Encoder的 zapcore.Core
func GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer, err := GetWriteSyncer(l.String()) // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return nil
	}
	return zapcore.NewCore(GetEncoder(), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(config.GetConfig().Zap.Prefix + t.Format("2006/01/02 - 15:04:05.000"))
}

// GetZapCores 根据配置文件的Level获取 []zapcore.Core
func GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	level := config.GetConfig().Zap.TransportLevel()

	for level := level; level <= zapcore.FatalLevel; level++ {
		cores = append(cores, GetEncoderCore(level, GetLevelPriority(level)))
	}

	return cores
}

// GetLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
func GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}

// GetWriteSyncer 获取 zapcore.WriteSyncer
func GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(config.GetConfig().Zap.Director, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(config.GetConfig().Zap.MaxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if config.GetConfig().Zap.LogInConsole {
		//return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), err
	}
	return zapcore.AddSync(fileWriter), err
}
