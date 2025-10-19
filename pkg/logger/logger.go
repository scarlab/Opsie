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
	requestFilePath = "/var/log/opsie/request.log"
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




func RequestLogger() (*log.Logger, *log.Logger) {
		requestFile := &lumberjack.Logger{
		Filename:   requestFilePath,
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     14,
		Compress:   true,
	}

	stdoutLogger := log.New(os.Stdout, bold+" HTTP  "+reset, log.Ldate|log.Ltime|log.Lmsgprefix)
	fileLogger := log.New(requestFile, " HTTP  ", log.Ldate|log.Ltime|log.Lmsgprefix)

	
	return  fileLogger, stdoutLogger
}