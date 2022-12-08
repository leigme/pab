package logger

import (
	"fmt"
	"github.com/leigme/pab/util"
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
	defaultLevel   = "info"
	logJsonFmt     = "json"
)

var logLevel = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

var sugaredLogger *zap.SugaredLogger

type Config struct {
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

func InitLoggerWithLogFile(filename string) {
	conf := Config{
		LogLevel:          defaultLevel,
		LogFormat:         "",
		LogPath:           filepath.Dir(filename),
		LogFileName:       filepath.Base(filename),
		LogFileMaxSize:    2,
		LogFileMaxBackups: 20,
		LogMaxAge:         7,
		LogCompress:       false,
		LogStdout:         false,
	}
	InitLoggerWithConfig(conf)
}

func InitLoggerWithConfig(conf Config) {
	encoder := getCatEncoder(conf.LogFormat)
	writeSyncer := getLogWriter(conf)
	level, ok := logLevel[strings.ToLower(conf.LogLevel)]
	if !ok {
		level = logLevel[defaultLevel]
	}
	InitLoggerWithCustom(encoder, writeSyncer, level)
}

func InitLoggerWithCustom(encoder zapcore.Encoder, writer zapcore.WriteSyncer, level zapcore.LevelEnabler) {
	core := zapcore.NewCore(encoder, writer, level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugaredLogger = logger.Sugar()
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
		encoder.AppendString(fmt.Sprintf("[%s] [_ti=%d] --", caller, util.GetThreadID()))
	}
	if strings.EqualFold(logFmt, logJsonFmt) {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriter 获取日志输出方式 日志文件 控制台
func getLogWriter(conf Config) zapcore.WriteSyncer {
	// 判断日志路径是否存在，如果不存在就创建
	if exist := util.IsExist(conf.LogPath); !exist {
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

func Debug(args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.DPanic(args)
}

func DPanicf(template string, args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.DPanicf(template, args...)
}

func Fatal(args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	defer sugaredLogger.Sync()
	sugaredLogger.Fatalf(template, args...)
}
