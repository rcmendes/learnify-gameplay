package logging

import (
	"log"
	"os"
)

var debugLogger = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
var infoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
var warningLogger = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
var errorLogger = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

//Debug is the log handler of DEBUG level of the application.
func Debug(messages ...interface{}) {
	debugLogger.Println(messages...)
}

//Info is the log handler of INFO level of the application.
func Info(messages ...interface{}) {
	infoLogger.Println(messages...)
}

//Warning is the log handler of WARNING level of the application.
func Warning(messages ...interface{}) {
	warningLogger.Println(messages...)
}

//Error is the log handler of ERROR level of the application.
func Error(messages ...interface{}) {
	errorLogger.Println(messages...)
}
