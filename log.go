package pab

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	defaultLogPath = "./"
	logJsonFmt     = "json"
)

var logLevel = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

type Logger struct {
	sugaredLogger *zap.SugaredLogger
}

type LoggerConfig struct {
	LogLevel          string `json:"log_level"`
	LogFormat         string `json:"log_format"`
	LogPath           string `json:"log_path"`
	LogFileName       string `json:"log_file_name"`
	LogFileMaxSize    int    `json:"log_file_max_size"` // 单位 mb
	LogFileMaxBackups int    `json:"log_file_max_backups"`
	LogMaxAge         int    `json:"log_max_age"` // 单位 day
	LogCompress       bool   `json:"log_compress"`
	LogStdout         bool   `json:"log_stdout"`
}

func NewLoggerWithCustom(encoder zapcore.Encoder, syncer zapcore.WriteSyncer, level zapcore.Level) Logger {
	core := zapcore.NewCore(encoder, syncer, level)
	logger := zap.New(core, zap.AddCaller())
	return Logger{
		sugaredLogger: logger.Sugar(),
	}
}

func NewLoggerWithConfig(conf LoggerConfig) Logger {
	encoder := getCatEncoder(conf.LogFormat)
	writeSyncer := getLogWriter(conf)
	level, ok := logLevel[strings.ToLower(conf.LogLevel)]
	if !ok {
		level = logLevel["info"]
	}
	return NewLoggerWithCustom(encoder, writeSyncer, level)
}

func NewLogger(filename, level string) Logger {
	conf := LoggerConfig{
		LogLevel:          level,
		LogFormat:         "",
		LogPath:           filepath.Dir(filename),
		LogFileName:       filepath.Base(filename),
		LogFileMaxSize:    2,
		LogFileMaxBackups: 20,
		LogMaxAge:         7,
		LogCompress:       false,
		LogStdout:         false,
	}
	return NewLoggerWithConfig(conf)
}

// isExist 判断文件或者目录是否存在
func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// getLogWriter 获取日志输出方式 日志文件 控制台
func getLogWriter(conf LoggerConfig) zapcore.WriteSyncer {
	// 判断日志路径是否存在，如果不存在就创建
	if exist := isExist(conf.LogPath); !exist {
		if strings.EqualFold(conf.LogPath, "") {
			conf.LogPath = defaultLogPath
		}
		err := os.MkdirAll(conf.LogPath, os.ModePerm)
		if err != nil {
			return nil
		}
	}
	// 日志文件 与 日志切割 配置
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(conf.LogPath, conf.LogFileName), // 日志文件路径
		MaxSize:    conf.LogFileMaxSize,                           // 单个日志文件最大多少 mb
		MaxBackups: conf.LogFileMaxBackups,                        // 日志备份数量
		MaxAge:     conf.LogMaxAge,                                // 日志最长保留时间
		Compress:   conf.LogCompress,                              // 是否压缩日志
	}

	if conf.LogStdout {
		// 日志同时输出到控制台和日志文件中
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout))
	} else {
		// 日志只输出到日志文件
		return zapcore.AddSync(lumberJackLogger)
	}
}

func getCatEncoder(logFmt string) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02T15:04:05.000"))
	}
	encoderConfig.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(fmt.Sprintf("[%s]", level.CapitalString()))
	}
	encoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(fmt.Sprintf("[%s] -- ", caller))
	}
	if strings.EqualFold(logFmt, logJsonFmt) {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
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
