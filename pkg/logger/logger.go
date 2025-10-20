package logger

import (
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

var fileLogger *log.Logger
var stdLogger *log.Logger

const (
	reset  = "\033[0m"
	bold = "\033[1m"

	// green  = "\033[32m"
	yellow = "\033[33m"
	red    = "\033[31m"
	cyan   = "\033[36m"
)


var (
	httpFilePath = "/var/log/opsie/http.log"
	logFilePath = "/var/log/opsie/opsie.log"
)


func Init() {

	logFile := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     14,
		Compress:   true,
	}



	stdLogger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmsgprefix)
	fileLogger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lmsgprefix)
}


func Printf(format string, v ...any) {
	stdLogger.Printf(format, v...)
	fileLogger.Printf(format, v...)
}

func Debug(format string, v ...any) {
	stdLogger.Printf(bold+cyan+" DEBUG "+reset+cyan+format+reset, v...)
	fileLogger.Printf(" DEBUG "+format, v...)
}

func Info(format string, v ...any) {
	// stdLogger.Printf(bold+green+" INFO  "+reset+green+format+reset, v...)
	stdLogger.Printf(" INFO  "+format, v...)
	fileLogger.Printf(" INFO  "+format, v...)
}

func Warn(format string, v ...any) {
	stdLogger.Printf(bold+yellow+" WARN  "+reset+yellow+format+reset, v...)
	fileLogger.Printf(" WARN  "+format, v...)
}

func Error(format string, v ...any) {
	stdLogger.Printf(bold+red+" ERROR "+reset+red+format+reset, v...)
	fileLogger.Printf(" ERROR "+format, v...)
}

func Fatalf(format string, v ...any) {
	stdLogger.Fatalf(bold+red+" FATAL_ERROR "+reset+red+format+reset, v...)
	fileLogger.Fatalf(" FATAL_ERROR "+format, v...)
}




func HttpLogger() (*log.Logger, *log.Logger) {
	httpFile := &lumberjack.Logger{
		Filename:   httpFilePath,
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     14,
		Compress:   true,
	}

	stdoutLogger := log.New(os.Stdout, bold+" HTTP  "+reset, log.Ldate|log.Ltime|log.Lmsgprefix)
	fileLogger := log.New(httpFile, " HTTP  ", log.Ldate|log.Ltime|log.Lmsgprefix)

	
	return  fileLogger, stdoutLogger
}