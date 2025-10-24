package logging

import (
	"fmt"
	"io"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger wraps zap logger
type Logger struct {
	*zap.Logger
}

// Config holds the logging configuration
type Config struct {
	Level            string   `json:"level"`
	Format           string   `json:"format"`
	OutputPaths      []string `json:"outputPaths"`
	ErrorOutputPaths []string `json:"errorOutputPaths"`
}

// NewLogger creates a new logger instance
func NewLogger(config Config) (*Logger, error) {
	// Set default values
	if config.Level == "" {
		config.Level = "info"
	}
	if config.Format == "" {
		config.Format = "json"
	}
	if len(config.OutputPaths) == 0 {
		config.OutputPaths = []string{"stdout"}
	}
	if len(config.ErrorOutputPaths) == 0 {
		config.ErrorOutputPaths = []string{"stderr"}
	}

	// Parse log level
	level := zapcore.InfoLevel
	if err := level.UnmarshalText([]byte(config.Level)); err != nil {
		return nil, fmt.Errorf("invalid log level: %v", err)
	}

	// Create encoder config
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Create encoder based on format
	var encoder zapcore.Encoder
	switch config.Format {
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	default:
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	// Create write syncers
	var outputs []zapcore.WriteSyncer
	for _, path := range config.OutputPaths {
		var writer io.Writer
		switch path {
		case "stdout":
			writer = os.Stdout
		case "stderr":
			writer = os.Stderr
		default:
			file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				return nil, fmt.Errorf("failed to open log file %s: %v", path, err)
			}
			writer = file
		}
		outputs = append(outputs, zapcore.AddSync(writer))
	}

	// Create error outputs
	var errorOutputs []zapcore.WriteSyncer
	for _, path := range config.ErrorOutputPaths {
		var writer io.Writer
		switch path {
		case "stdout":
			writer = os.Stdout
		case "stderr":
			writer = os.Stderr
		default:
			file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				return nil, fmt.Errorf("failed to open error log file %s: %v", path, err)
			}
			writer = file
		}
		errorOutputs = append(errorOutputs, zapcore.AddSync(writer))
	}

	// Create core
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(outputs...),
		level,
	)

	// Create logger
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return &Logger{logger}, nil
}

// WithFields adds fields to the logger
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return &Logger{l.Logger.With(zapFields...)}
}

// WithField adds a single field to the logger
func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{l.Logger.With(zap.Any(key, value))}
}

// Close closes the logger and flushes any buffered log entries
func (l *Logger) Close() error {
	return l.Logger.Sync()
}

// DefaultLogger returns a default logger configuration
func DefaultLogger() *Logger {
	config := Config{
		Level:            "info",
		Format:           "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := NewLogger(config)
	if err != nil {
		// Fallback to standard logger if zap logger fails
		log.Printf("Failed to create zap logger: %v", err)
		return &Logger{zap.NewNop()}
	}

	return logger
}

// DevelopmentLogger returns a logger configured for development
func DevelopmentLogger() *Logger {
	config := Config{
		Level:            "debug",
		Format:           "console",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := NewLogger(config)
	if err != nil {
		// Fallback to standard logger if zap logger fails
		log.Printf("Failed to create zap logger: %v", err)
		return &Logger{zap.NewNop()}
	}

	return logger
}
