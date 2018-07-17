package logger

import "go.uber.org/zap"

// New creates package specific loging pipeline.
func New(name string) *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // nolint:errcheck
	return logger.Named(name).Sugar()
}
