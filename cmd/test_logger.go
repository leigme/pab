package main

import (
	"github.com/leigme/pab/logger"
	"time"
)

func init() {
	logger.InitLoggerWithLogFile("loggerDemo.log")
}

func main() {
	go func() {
		logger.Info("协程日志")
	}()
	for i := 0; i < 10; i++ {
		logger.Infof("主线程日志: %d", i*10)
		time.Sleep(time.Duration(3) * time.Second)
	}
}
