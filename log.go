package pab

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"time"
)

func zapTest() {
	logger := NewLoggerWithCatInfoAndError("./demo_info.log", "./demo_error.log")
	logger.Infof("info 日志 %s", "abc")
	logger.Errorf("error 日志 %d", 30)
	logger.Errorf("error 日志 %d", 30)
	logger.Errorf("error 日志 %d", 30)
	logger.Errorf("error 日志 %d", 30)
	logger.Errorf("error 日志 %d", 30)
	logger.Errorf("error 日志 %d", 30)
	logger.Errorf("error 日志 %d", 30)
	logger.Infof("info 日志 %s", "abc")
	logger.Infof("info 日志 %s", "abc")
	logger.Infof("info 日志 %s", "abc")
}

type Logger struct {
	sugaredLogger *zap.SugaredLogger
	logFile       string
}

type LoggerFile struct {
	Encoder   zapcore.Encoder
	LogWriter io.Writer
	Level     zap.LevelEnablerFunc
}

func encoderWithCat() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
		EncodeLevel: func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(strings.ToUpper(fmt.Sprintf("[%s]", level)))
		},
		TimeKey: "ts",
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format("2006-01-02T15:04:05"))
		},
		CallerKey: "false",
		EncodeCaller: func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(fmt.Sprintf("[%s] -- ", caller))
		},
		EncodeDuration: func(duration time.Duration, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendInt64(int64(duration) / 1000000)
		},
	})
}

func getWriter(filename string) io.Writer {
	logFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return logFile
}

func createCores(loggerFiles ...LoggerFile) []zapcore.Core {
	cores := make([]zapcore.Core, 0)
	for _, loggerFile := range loggerFiles {
		cores = append(cores, zapcore.NewCore(loggerFile.Encoder, zapcore.AddSync(loggerFile.LogWriter), loggerFile.Level))
	}
	return cores
}

func NewLoggerWithCat(logFile string) Logger {
	infoLoggerFile := LoggerFile{
		encoderWithCat(),
		getWriter(logFile),
		zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level >= zapcore.InfoLevel
		}),
	}
	return NewLogger(createCores(infoLoggerFile)...)
}

func NewLogger(cores ...zapcore.Core) Logger {
	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller())
	return Logger{
		sugaredLogger: logger.Sugar(),
	}
}

func NewLoggerWithCatInfoAndError(infoLogFile, errorLogFile string) Logger {
	infoLoggerFile := LoggerFile{
		encoderWithCat(),
		getWriter(infoLogFile),
		zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level >= zapcore.InfoLevel
		}),
	}
	errorLoggerFile := LoggerFile{
		encoderWithCat(),
		getWriter(errorLogFile),
		zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level >= zapcore.ErrorLevel
		}),
	}
	cores := createCores(infoLoggerFile, errorLoggerFile)
	return NewLogger(cores...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.sugaredLogger.Debugf(template, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.sugaredLogger.Infof(template, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.sugaredLogger.Warnf(template, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.sugaredLogger.Errorf(template, args...)
}

func (l *Logger) DPanic(args ...interface{}) {
	l.sugaredLogger.DPanic(args...)
}

func (l *Logger) DPanicf(template string, args ...interface{}) {
	l.sugaredLogger.DPanicf(template, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.sugaredLogger.Fatalf(template, args...)
}
