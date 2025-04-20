// Package that give basic structure to the standard library
// https://pkg.go.dev/log package. This package allows to define logging
// level at which the logging messages will be print to the standard output.
package go_logger

import (
	"log"
	"os"
)

// Log levels
const(
	InfoLevel = iota // Logging level representing Information level
	WarningLevel // Logging level representing Warning level
	ErrorLevel // Logging level representing Error level
)

// Struct representing the logger structure.
type Logger struct {
	Level		int
	infoLogger	*log.Logger
	warnLogger	*log.Logger
	errorLogger	*log.Logger
}

var logger *Logger

// init function
func init() {
	// Initializer method for the logging level and the log objects (logger)for each level.
	logger = &Logger{
		Level: InfoLevel, // default value
		infoLogger: log.New(os.Stdout, "INFO: ", log.LstdFlags),
		warnLogger: log.New(os.Stdout, "WARN: ", log.LstdFlags),
		errorLogger: log.New(os.Stdout, "ERROR: ", log.LstdFlags | log.Lshortfile),
	}
}

// Method to set the logging level.
func SetLevel(level int) {
	logger.Level = level
}

// Methods to log (at different levels)

// Method to print information messages
func Info(message string) {
	if logger.Level <= InfoLevel {
		logger.infoLogger.Println(message)
	}
}

// Method to print warning messages
func Warning(message string) {
	if logger.Level <= WarningLevel {
		logger.warnLogger.Println(message)
	}
}

// Method to print error messages
func Error(message string) {
	if logger.Level <= ErrorLevel {
		logger.errorLogger.Println(message)
	}
}

