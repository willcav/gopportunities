package config

func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	// Intialize Logger
	logger := NewLogger(p)
	return logger
}
