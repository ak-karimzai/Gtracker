// Package logger provides an interface for logging functionality.
package logger

// Logger represents a logger interface with various logging methods.
type Logger interface {
	// Fatal logs a message at fatal level and exits the application.
	Fatal(v ...interface{})

	// Fatalf logs a formatted message at fatal level and exits the application.
	Fatalf(format string, v ...interface{})

	// Panic logs a message at panic level and panics.
	Panic(v ...interface{})

	// Panicf logs a formatted message at panic level and panics.
	Panicf(format string, v ...interface{})

	// Print logs a message at default level.
	Print(v ...interface{})

	// Printf logs a formatted message at default level.
	Printf(format string, v ...interface{})

	// Println logs a message at default level followed by a newline.
	Println(v ...interface{})

	// Error logs a message at error level.
	Error(v ...interface{})

	// Errorf logs a formatted message at error level.
	Errorf(format string, v ...interface{})

	// Info logs informational messages.
	Info(args ...interface{})
}
