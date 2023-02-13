package logging

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
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

func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = openLogFile(fileName, filePath)
	if err != nil {
		log.Fatalln(err)
	}
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...any) {
	setPrefix(DEBUG)
	logger.Println(time.Now().UnixNano()/1e6, ": ", v)
}

func Info(v ...any) {
	setPrefix(INFO)
	logger.Println(time.Now().UnixNano()/1e6, ": ", v)
	fmt.Println(time.Now().UnixNano()/1e6, ": ", v)
}

func Warn(v ...any) {
	setPrefix(WARNING)
	logger.Println(time.Now().UnixNano()/1e6, ": ", v)
}

func Error(v ...any) {
	setPrefix(ERROR)
	logger.Println(time.Now().UnixNano()/1e6, ": ", v)
}

func Fatal(v ...any) {
	setPrefix(FATAL)
	logger.Println(time.Now().UnixNano()/1e6, ": ", v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], file, line)
	} else {
		logPrefix = fmt.Sprintf("%s", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
