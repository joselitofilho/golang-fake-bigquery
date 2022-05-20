package logging

import (
	"log"
	"os"
)

var (
	Debug   = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	Trace   = log.New(os.Stdout, "[TRACE] ", log.Ldate|log.Ltime|log.Lshortfile)
	Info    = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stderr, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	Error   = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
)
