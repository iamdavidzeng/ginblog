package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
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
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.Lshortfile)
}

func Debug(v ...interface{}) {
	setPreifx(DEBUG)
	logger.Println(v)

}

func Info(v ...interface{}) {
	setPreifx(INFO)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPreifx(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPreifx(FATAL)
	logger.Println(v)
}

func setPreifx(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
