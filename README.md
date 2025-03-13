# ZLogger

![Go Version](https://img.shields.io/badge/Go-%3E%3D1.17-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![Build](https://github.com/pablolagos/zlogger/workflows/Go/badge.svg)

**ZLogger** is a lightweight logging library built on top of [`zerolog`](https://github.com/rs/zerolog) with support for automatic log rotation and optional integration with Sentry.

## Features

- **Fast and efficient** logging using `zerolog`.
- **Automatic log rotation** with [`lumberjack`](https://github.com/natefinch/lumberjack).
- **Customizable log levels** with colored output.
- **Sentry integration** for error tracking.
- **Minimal setup required**.

## Installation

```sh
 go get github.com/pablolagos/zlogger
```

## Usage

### Basic Logger

```go
package main

import (
	"github.com/pablolagos/zlogger"
)

func main() {
	logger := zlogger.New("app.log", 10, 3) // Log file with rotation
	logger.Info("Application started")
	logger.Warn("This is a warning message")
	logger.Error("An error occurred")
}
```

### Standard Error Logger

```go
logger := zlogger.NewStdErr()
logger.Debug("Debugging mode enabled")
```

### Logger with Sentry Integration

```go
logger := zlogger.NewWithSentry("app.log", 10, 3, "your_sentry_dsn", "v1.0.0", "production")
logger.Error("This error will be reported to Sentry")
```

## API Reference

### `New(filename string, maxSize int, maxBackups int) *ZLogger`
Creates a new logger with automatic log rotation.

- `filename`: Path to the log file. If empty, logs are written to `stderr`.
- `maxSize`: Maximum log file size in MB before rotation.
- `maxBackups`: Number of backup logs to retain.

### `NewStdErr() *ZLogger`
Creates a new logger that writes logs to `stderr`.

### `NewWithSentry(filename string, maxSize int, maxBackups int, dsn string, release string, environment string) *ZLogger`
Creates a logger with Sentry integration.

- `dsn`: Sentry DSN for error tracking.
- `release`: Application release version.
- `environment`: Deployment environment (`production`, `staging`, etc.).

### Logging Methods

| Method     | Description  |
|------------|-------------|
| `Debug(v ...interface{})` | Logs a debug message |
| `Info(v ...interface{})` | Logs an info message |
| `Warn(v ...interface{})` | Logs a warning |
| `Error(v ...interface{})` | Logs an error |
| `Fatal(v ...interface{})` | Logs a fatal error and exits |
| `Panic(v ...interface{})` | Logs a panic message and panics |
| `Trace(v ...interface{})` | Logs a trace message |
| `Debugf(format string, v ...interface{})` | Logs a formatted debug message |
| `Infof(format string, v ...interface{})` | Logs a formatted info message |
| `Warnf(format string, v ...interface{})` | Logs a formatted warning message |
| `Errorf(format string, v ...interface{})` | Logs a formatted error message |
| `Fatalf(format string, v ...interface{})` | Logs a formatted fatal error and exits |
| `Panicf(format string, v ...interface{})` | Logs a formatted panic message and panics |


## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License.

## Author

Developed by [Pablo Lagos](https://github.com/pablolagos).

---

⭐ If you like this project, don't forget to star it on GitHub!

