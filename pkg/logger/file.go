package logger

import (
	"fmt"
	"go-talk-talk/config"
	"log"
	"os"
	"time"
)

var LogTimeFormat = "20060102"

func openLogFile() *os.File {
	fileName := fmt.Sprintf("%s%s.%s", config.LoggerConf.LogSaveName, time.Now().Format(LogTimeFormat), config.LoggerConf.LogFileExt)
	filePath := config.LoggerConf.LogPath + fileName
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission: %v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}
	return handle
}

func mkDir() {
	// 获取当前项目目录地址
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+config.LoggerConf.LogPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
