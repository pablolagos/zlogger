// Package zlogger is a wrapper for zerolog with specific format and optionally integrated with Sentry
package zlogger

import (
	"fmt"
	"io"
	stdlog "log"
	"os"
	"path"
	"strings"

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

// New creates a console zerolog with auto rotating feature
//
//	Filename: Filename to write log. If empty, stderr will be used.
//	MaxSize: Max size before rotating, in MB
//	MaxBackups: Number of backups to retain. 0=unlimited
//
//nolint:cyclop // This function is not too complex
func New(filename string, maxSize int, maxBackups int) *ZLogger {
	var logger zerolog.Logger
	var outFile io.Writer
	z := ZLogger{}
	z.setLevelNames(true)

	if len(filename) > 0 {
		outFile = &lumberjack.Logger{
			Filename:   path.Join(filename),
			MaxSize:    maxSize,
			MaxBackups: maxBackups,
			Compress:   true, // disabled by default
		}
	} else {
		outFile = os.Stderr
	}

	logger = zerolog.New(zerolog.ConsoleWriter{
		Out:        outFile,
		TimeFormat: "2006/02/01 15:04:05",
		FormatLevel: func(i interface{}) string {
			switch i {
			case zerolog.LevelErrorValue:
				return z.levelError
			case zerolog.LevelFatalValue:
				return z.levelFatal
			case zerolog.LevelWarnValue:
				return z.levelWarning
			case zerolog.LevelInfoValue:
				return z.levelInfo
			case zerolog.LevelDebugValue:
				return z.levelDebug
			case zerolog.LevelTraceValue:
				return z.levelTrace
			case zerolog.LevelPanicValue:
				return z.levelPanic
			case nil:
				return emptyString
			default:
				return strings.ToUpper(fmt.Sprintf("[%s]", i))
			}
		},
	}).With().Timestamp().Logger()
	z.Logger = logger

	return &z
}

// NewStdErr creates a zerolog with stderr output, for testing purposes
func NewStdErr() *ZLogger {
	return New("", 0, 0)
}

// NewWithSentry creates a zerolog with auto rotating feature and Sentry integration
//
//	Filename: Filename to write log. If empty, stderr will be used.
//	MaxSize: Max size before rotating, in MB
//	MaxBackups: Number of backups to retain. 0=unlimited
func NewWithSentry(filename string, maxSize int, maxBackups int, dsn, release, environment string) *ZLogger {
	var logger zerolog.Logger
	var outFile io.Writer
	z := ZLogger{}
	z.setLevelNames(true)

	if len(filename) > 0 {
		outFile = &lumberjack.Logger{
			Filename:   path.Join(filename),
			MaxSize:    maxSize,
			MaxBackups: maxBackups,
			Compress:   true, // disabled by default
		}
	} else {
		outFile = os.Stderr
	}

	writer1 := zerolog.ConsoleWriter{
		Out:        outFile,
		TimeFormat: "2006/02/01 15:04:05",
		FormatLevel: func(i interface{}) string {
			switch i {
			case zerolog.LevelErrorValue:
				return z.levelError
			case zerolog.LevelFatalValue:
				return z.levelFatal
			case zerolog.LevelWarnValue:
				return z.levelWarning
			case zerolog.LevelInfoValue:
				return z.levelInfo
			case zerolog.LevelDebugValue:
				return z.levelDebug
			case zerolog.LevelTraceValue:
				return z.levelTrace
			case zerolog.LevelPanicValue:
				return z.levelPanic
			case nil:
				return emptyString
			default:
				return strings.ToUpper(fmt.Sprintf("[%s]", i))
			}
		},
	}

	// Sentry integration
	scope := sentry.NewScope()
	client, _ := sentry.NewClient(sentry.ClientOptions{
		Dsn:              dsn,
		Debug:            true,
		AttachStacktrace: true,
		Release:          release,
		Environment:      environment,
	})
	_ = sentry.NewHub(client, scope)

	w, err := zlogsentry.New(client)
	if err != nil {
		stdlog.Fatal(err)
	}
	logger = zerolog.New(io.MultiWriter(w, writer1)).With().Timestamp().Logger()
	z.Logger = logger

	return &z
}

func (z *ZLogger) setLevelNames(applyColors bool) {
	color.NoColor = !applyColors

	if !applyColors {
		z.levelTrace = "[TRACE]"
		z.levelInfo = "[INFO]"
		z.levelWarning = "[WARN]"
		z.levelError = "[ERROR]"
		z.levelFatal = "[FATAL]"
		z.levelDebug = "[DEBUG]"
		z.levelPanic = "[PANIC]"

		return
	}

	z.levelTrace = "[TRACE]"
	z.levelInfo = color.BlueString("[INFO]")
	z.levelWarning = color.YellowString("[WARN]")
	z.levelError = color.RedString("[ERROR]")
	z.levelFatal = color.New(color.BgRed).Add(color.FgWhite).Sprint("[FATAL]")
	z.levelDebug = color.HiBlueString("[DEBUG]")
	z.levelPanic = color.New(color.BgHiRed).Add(color.BgBlack).Sprint("[PANIC]")
}

func (z *ZLogger) GetLogger() zerolog.Logger {
	return z.Logger
}

func (z *ZLogger) Debug(v ...interface{}) {
	z.Logger.Debug().Msg(fmt.Sprint(v...))
}

func (z *ZLogger) Info(v ...interface{}) {
	z.Logger.Info().Msg(fmt.Sprint(v...))
}

func (z *ZLogger) Error(v ...interface{}) {
	z.Logger.Error().Msg(fmt.Sprint(v...))
}

func (z *ZLogger) Warn(v ...interface{}) {
	z.Logger.Warn().Msg(fmt.Sprint(v...))
}

func (z *ZLogger) Debugf(format string, v ...interface{}) {
	z.Logger.Debug().Msgf(format, v...)
}

func (z *ZLogger) Infof(format string, v ...interface{}) {
	z.Logger.Info().Msgf(format, v...)
}

func (z *ZLogger) Errorf(format string, v ...interface{}) {
	z.Logger.Error().Msgf(format, v...)
}

func (z *ZLogger) Warnf(format string, v ...interface{}) {
	z.Logger.Warn().Msgf(format, v...)
}
