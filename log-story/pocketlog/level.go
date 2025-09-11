package pocketlog

// Level represent an available logging level
type Level byte

const (
	// LevelDebug represents the lowest level of log, mostly used
	// debugging purposes.
	LevelDebug Level = iota

	// LevelInfo represents a logging level that contains information
	// deemed valuable.
	LevelInfo

	// LevelError represents the highest logging level, only to be used
	// to trace errors.
	LevelError
)
