package logger

// Logger defines a minimal logging interface. Different logger implementations
// (for example the JSON logger in internal/infrastructure/logger/jsonlog)
// should satisfy this interface so they can be swapped easily.
type Logger interface {
	PrintInfo(message string, properties map[string]string)
	PrintError(message string, properties map[string]string)
	PrintFatal(message string, properties map[string]string)
	Write(message []byte) (int, error)
}
