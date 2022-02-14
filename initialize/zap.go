package initialize

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"server/utils"
	"time"
)

var path = "log"

func InitLogger() *zap.Logger {
	if ok, _ := utils.PathExists(path); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", path)
		_ = os.Mkdir(path, os.ModePerm)
	}
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/debug.log", path), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/info.log", path), infoPriority),
		getEncoderCore(fmt.Sprintf("./%s/warn.log", path), warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/error.log", path), errorPriority),
	}

	return zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())
}

func getEncoderConfig() (encoderConfig zapcore.EncoderConfig) {
	encoderConfig = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	return
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[asset-management-system]" + "2006/01/02 - 15:04:05.000"))
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writeSyncer := getWriter(fileName)
	return zapcore.NewCore(getEncoder(), writeSyncer, level)
}

func getWriter(fileName string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
