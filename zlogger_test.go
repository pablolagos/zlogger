package zlogger

import "testing"

func TestAll(t *testing.T) {
	// Initialize logger
	logger := NewStdErr()

	logger.Info("Info message")
	logger.Debug("Debug message")
	logger.Error("Error message")
	logger.Warn("Warning message")
	logger.Debugf("Debug message with format: %s", "formatted")
	logger.Infof("Info message with format: %s", "formatted")
	logger.Errorf("Error message with format: %s", "formatted")
	logger.Warnf("Warning message with format: %s", "formatted")
	logger.Printf("Print message with format: %s", "formatted")
	logger.Println("Println message", "field1", "field2")
}
