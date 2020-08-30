package logging

import (
	"fmt"
	"log"
	"os"
)

type Level int

var (
	DefaultPrefix      = "[iamdavidzeng]"
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	logger = log.New(os.Stdout, DefaultPrefix, log.Ldate|log.Ltime|log.Lshortfile)
}

// Debug report
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

// Info report
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

// Warn report
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

// Error report
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

// Fatal report
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	logPrefix = fmt.Sprintf("[%s]: ", levelFlags[level])
	logger.SetPrefix(logPrefix)
}
