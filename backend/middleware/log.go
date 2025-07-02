package middleware

import (
	"log"
	"os"
)

var (
	WarningLog *log.Logger
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
)

// TODO: do a better log, maybe use slog?
func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		log.Fatalf("failed to open logs.txt file: %v", err)
	}

	InfoLog = log.New(file, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLog = log.New(file, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(file, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}
