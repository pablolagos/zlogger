package zlogger

import (
	"context"
	"io"
	"os"
	"path"
	"time"

	"github.com/fatih/color"
	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	zlogsentry "github.com/sveatlo/zerolog-sentry"
	"gopkg.in/natefinch/lumberjack.v2"
)

const emptyString = ""

type ZLogger struct {
	zerolog.Logger
	levelTrace   string
	levelInfo    string
	levelWarning string
	levelError   string
	levelFatal   string
	levelDebug   string
	levelPanic   string
}

// newLogger creates a zerolog instance with optional file rotation
func newLogger(filename string, maxSize, maxBackups int, enableColors bool) (zerolog.Logger, io.Writer) {
	var outFile io.Writer

	if filename != "" {
		outFile = &lumberjack.Logger{
			Filename:   path.Join(filename),
			MaxSize:    maxSize,
			MaxBackups: maxBackups,
			Compress:   true,
		}
	} else {
		outFile = os.Stderr
	}

	logger := zerolog.New(zerolog.ConsoleWriter{
		Out:        outFile,
		TimeFormat: "2006/02/01 15:04:05",
	}).With().Timestamp().Logger()

	return logger, outFile
}

// New creates a zerolog instance with optional file rotation
func New(filename string, maxSize, maxBackups int, enableColors bool) *ZLogger {
	logger, _ := newLogger(filename, maxSize, maxBackups, enableColors)
	z := &ZLogger{Logger: logger}
	z.setLevelNames(enableColors)
	return z
}

// NewStdErr creates a zerolog instance that logs to stderr (for testing purposes)
func NewStdErr() *ZLogger {
	return New("", 0, 0, true) // Logs to stderr with default settings
}

// NewWithSentry creates a zerolog instance with Sentry integration
func NewWithSentry(filename string, maxSize, maxBackups int, enableColors bool, dsn, release, environment string) *ZLogger {
	logger, outFile := newLogger(filename, maxSize, maxBackups, enableColors)
	z := &ZLogger{Logger: logger}
	z.setLevelNames(enableColors)

	scope := sentry.NewScope()
	client, err := sentry.NewClient(sentry.ClientOptions{
		Dsn:              dsn,
		Debug:            true,
		AttachStacktrace: true,
		Release:          release,
		Environment:      environment,
	})
	if err != nil {
		z.Logger.Warn().Msg("Failed to initialize Sentry logging")
		return z
	}
	_ = sentry.NewHub(client, scope)
	w, err := zlogsentry.New(client)
	if err != nil {
		z.Logger.Warn().Msg("Failed to create Sentry writer")
		return z
	}

	z.Logger = zerolog.New(io.MultiWriter(w, outFile)).With().Timestamp().Logger()
	defer sentry.Flush(2 * time.Second)
	return z
}

// setLevelNames configures level names with or without colors
func (z *ZLogger) setLevelNames(applyColors bool) {
	color.NoColor = !applyColors

	z.levelTrace = "[TRACE]"
	z.levelInfo = color.BlueString("[INFO]")
	z.levelWarning = color.YellowString("[WARN]")
	z.levelError = color.RedString("[ERROR]")
	z.levelFatal = color.New(color.BgRed).Add(color.FgWhite).Sprint("[FATAL]")
	z.levelDebug = color.HiBlueString("[DEBUG]")
	z.levelPanic = color.New(color.BgHiRed).Add(color.BgBlack).Sprint("[PANIC]")
}

// GetLogger returns the underlying zerolog instance
func (z *ZLogger) GetLogger() zerolog.Logger {
	return z.Logger
}

// Logging functions without formatting
func (z *ZLogger) Debug(v ...interface{}) { z.Logger.Debug().Interface("data", v).Msg("") }
func (z *ZLogger) Info(v ...interface{})  { z.Logger.Info().Interface("data", v).Msg("") }
func (z *ZLogger) Error(v ...interface{}) { z.Logger.Error().Interface("data", v).Msg("") }
func (z *ZLogger) Warn(v ...interface{})  { z.Logger.Warn().Interface("data", v).Msg("") }

// Logging functions with formatting
func (z *ZLogger) Debugf(format string, v ...interface{}) { z.Logger.Debug().Msgf(format, v...) }
func (z *ZLogger) Infof(format string, v ...interface{})  { z.Logger.Info().Msgf(format, v...) }
func (z *ZLogger) Errorf(format string, v ...interface{}) { z.Logger.Error().Msgf(format, v...) }
func (z *ZLogger) Warnf(format string, v ...interface{})  { z.Logger.Warn().Msgf(format, v...) }

// Context-aware logging functions
func (z *ZLogger) DebugCtx(ctx context.Context, v ...interface{}) {
	z.Logger.Debug().Ctx(ctx).Interface("data", v).Msg("")
}
func (z *ZLogger) InfoCtx(ctx context.Context, v ...interface{}) {
	z.Logger.Info().Ctx(ctx).Interface("data", v).Msg("")
}
func (z *ZLogger) ErrorCtx(ctx context.Context, v ...interface{}) {
	z.Logger.Error().Ctx(ctx).Interface("data", v).Msg("")
}
func (z *ZLogger) WarnCtx(ctx context.Context, v ...interface{}) {
	z.Logger.Warn().Ctx(ctx).Interface("data", v).Msg("")
}
