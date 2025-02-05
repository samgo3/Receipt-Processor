package utils

import (
	"log"
	"sync"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger     *zap.Logger
	loggerOnce sync.Once
)

func initializeLogger() {
	config := zap.NewProductionConfig()
	config.DisableStacktrace = true
	logLevel := viper.GetString("log.level")
	level, err := zapcore.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("Invalid log level in config: %v", logLevel)
	}
	config.Level = zap.NewAtomicLevelAt(level)
	logger = zap.Must(config.Build())
	defer logger.Sync()
}

// GetLogger returns a singleton instance of the logger.
func GetLogger() *zap.Logger {
	loggerOnce.Do(
		initializeLogger)
	return logger
}
