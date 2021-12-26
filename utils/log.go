package utils

import (
	"BackendScreenShot/constants"
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

// log
// TODO
// log level 支持
// 异步写文件
// 要用buf

var log *EasyLog
var logOnce sync.Once

type EasyLog struct {
	//	logPath string
	level   EasyLogLevel
	writter *bufio.Writer
}

type EasyLogLevel int

const (
	EasyLogDebug EasyLogLevel = iota
	EasyLogInfo
	EasyLogWarn
	EasyLogError
)

func (l *EasyLog) Debug(msg ...interface{}) {
	if l.level > EasyLogDebug {
		return
	}
	l.writter.Write([]byte(l.formatMsg(msg)))
}

func (l *EasyLog) Info(msg ...interface{}) {
	fmt.Println(l.level, EasyLogInfo)
	if l.level > EasyLogInfo {
		return
	}

	l.writter.Write([]byte(l.formatMsg(msg)))
}

func (l *EasyLog) Warn(msg ...interface{}) {
	if l.level > EasyLogWarn {
		return
	}
	l.writter.Write([]byte(l.formatMsg(msg)))
}

func (l *EasyLog) Error(msg ...interface{}) {
	if l.level > EasyLogError {
		return
	}
	l.writter.Write([]byte(l.formatMsg(msg)))
}

func (l *EasyLog) formatMsg(msg ...interface{}) string {
	var str string
	for _, v := range msg {
		str += fmt.Sprintf("% v", v)
	}
	return time.Now().Format(time.RFC3339) + str + "\n"
}

func (l *EasyLog) Clean() {
	l.writter.Flush()
}

func Log() *EasyLog {
	logOnce.Do(func() {
		// open log file
		file, err := os.OpenFile(constants.LogPath+constants.ServictName+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0b111111111)
		if err != nil {
			panic(err)
		}
		// defer file.Close()
		log = &EasyLog{
			level:   EasyLogDebug,
			writter: bufio.NewWriter(file),
		}
	})
	return log
}
