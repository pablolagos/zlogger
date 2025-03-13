# ZLogger - Advanced Structured Logging for Go

ZLogger is a structured logging library for Go, built on top of [zerolog](https://github.com/rs/zerolog). It provides an easy-to-use API with advanced features like:
- **Log rotation** with [lumberjack](https://github.com/natefinch/lumberjack)
- **Sentry integration** for automatic error tracking
- **Custom log levels with colors** for better readability
- **Context-aware logging** for structured and contextual debugging

## 🚀 Features
- 🔥 **High-performance structured logging** using zerolog.
- 📁 **Automatic log rotation** to prevent log files from growing indefinitely.
- 🛠 **Sentry integration** to capture and monitor errors seamlessly.
- 🎨 **Customizable log level names and colors**.
- 📡 **Context-aware logging** to enrich log data.
- 🌍 **Lightweight and efficient** with minimal overhead.

---

## 📦 Installation
```sh
# Install ZLogger using go modules
 go get github.com/pablolagos/zlogger
```

---

## ⚡ Usage

### **Basic Logging**
```go
package main

import (
	"github.com/pablolagos/zlogger"
)

func main() {
	logger := zlogger.New("app.log", 10, 5, true) // 10MB max size, 5 backups, colors enabled
	logger.Info("Application started")
	logger.Debug("Debugging details...")
	logger.Error("An error occurred")
}
```

### **Logging with Sentry**
```go
logger := zlogger.NewWithSentry("app.log", 10, 5, true, "your_sentry_dsn", "1.0.0", "production")
logger.Error("Critical error: Database connection failed")
```

### **Context-aware Logging**
```go
import "context"

ctx := context.WithValue(context.Background(), "request_id", "12345")
logger.InfoCtx(ctx, "Processing request")
```

### **Logging to Stderr (Testing)**
```go
logger := zlogger.NewStdErr()
logger.Info("This message logs to stderr")
```

---

## 🔔 How Sentry Works with ZLogger
ZLogger integrates seamlessly with [Sentry](https://sentry.io/) to capture and monitor errors in your Go application. When using `NewWithSentry()`, ZLogger automatically sends logs of level `Error` and above to Sentry.

### **How it Works**
1. ZLogger initializes a Sentry client with the provided DSN (Data Source Name).
2. Errors, warnings, or fatal logs are captured and sent to Sentry.
3. Sentry records logs with stack traces and metadata (like environment and release version).
4. You can view and analyze errors in your Sentry dashboard.

### **Example Configuration**
```go
logger := zlogger.NewWithSentry("app.log", 10, 5, true, "your_sentry_dsn", "1.0.0", "production")
logger.Error("Database connection failed")
```

### **Sentry Best Practices**
- Ensure your **DSN is correctly configured** in environment variables.
- Use meaningful **release versions** to track issues across deployments.
- Call `sentry.Flush(time.Second * 2)` before exiting the application to ensure logs are sent.

---

## 🎨 Log Level Customization
ZLogger supports custom log level names and colors:
- `INFO`: Blue
- `WARN`: Yellow
- `ERROR`: Red
- `FATAL`: Red background with white text
- `DEBUG`: High-intensity blue

To disable colors, pass `false` in the `New()` function:
```go
logger := zlogger.New("app.log", 10, 5, false) // Colors disabled
```

---

## 🛠 Available Methods
ZLogger provides the following logging methods:

### **Standard Logging Methods**
```go
logger.Debug("Debug message")
logger.Info("Informational message")
logger.Warn("Warning message")
logger.Error("Error message")
```

### **Formatted Logging Methods**
```go
logger.Debugf("Debugging: %s", "details")
logger.Infof("User %s logged in", "John")
logger.Warnf("Warning: %d attempts detected", 3)
logger.Errorf("Error: %v", err)
```

### **Context-aware Logging Methods**
```go
logger.DebugCtx(ctx, "Debug message with context")
logger.InfoCtx(ctx, "Info message with context")
logger.WarnCtx(ctx, "Warning message with context")
logger.ErrorCtx(ctx, "Error message with context")
```

---

## 🛠 Configuration Options
| Parameter    | Type    | Description |
|-------------|--------|-------------|
| `filename`  | string | Log file path (empty to use stderr) |
| `maxSize`   | int    | Max log file size in MB before rotation |
| `maxBackups`| int    | Number of rotated logs to retain (0 = unlimited) |
| `enableColors` | bool | Enable/disable color output |

---

## 🔥 Why Use ZLogger?
- **Performance:** Efficient structured logging with low memory overhead.
- **Flexibility:** Works with stdout, file-based logging, and remote monitoring (Sentry).
- **Simplicity:** Easy-to-use API with sane defaults.
- **Scalability:** Suitable for microservices, monoliths, and cloud-based applications.

---

## 🛡 License
This project is licensed under the MIT License.

---

## 👨‍💻 Contributing
We welcome contributions! Feel free to submit issues and pull requests to improve ZLogger.

---


## Author

Developed by [Pablo Lagos](https://github.com/pablolagos).

---

⭐ If you like this project, don't forget to star it on GitHub!

---

**Happy Logging! 🚀**
