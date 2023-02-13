package logging

import (
	"fmt"
	"github.com/JQZhangC/framework/pkg/file"
	"os"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", "runtime/", "logs/")
}

func getLogFileName() string {
	return fmt.Sprintf("log%s.%s",
		time.Now().Format("20060102"),
		"log")
}

func openLogFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file."+
			"CheckPermission Permission denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}
	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("fail to OpenFile :%v", err)
	}
	return f, nil
}
