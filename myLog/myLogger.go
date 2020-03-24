package myLog

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type logLevel uint

//构造函数
type myLogger struct {
	logLevel
}

const (
	UNKNOWN logLevel = iota
	DEBUG
	TRACE
	INFO
	WARRING
	ERROR
	FATAL
)

func parseLevel(level string) (logLevel, error) {
	switch strings.ToLower(level) {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil

	case "info":
		return INFO, nil

	case "warning":
		return WARRING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		return UNKNOWN, errors.New("不是合法的日志级别")
	}
}

//noinspection GoExportedFuncWithUnexportedType
func NewLogger(level string) (myLogger, error) {

	ll, err := parseLevel(level)
	if err != nil {
		fmt.Println(err)
		return myLogger{UNKNOWN}, err
	}
	log := myLogger{
		ll,
	}
	return log, nil
}
func judge(logLevel, level logLevel) bool {
	return logLevel <= level
}
func (log myLogger) Debug(format string, a ...interface{}) {
	logPri(log.logLevel, DEBUG, format, a)

}

func (log myLogger) Trace(format string, a ...interface{}) {
	logPri(log.logLevel, TRACE, format, a)

}

func (log myLogger) Info(format string, a ...interface{}) {
	logPri(log.logLevel, INFO, format, a)

}

func (log myLogger) Warning(format string, a ...interface{}) {
	logPri(log.logLevel, WARRING, format, a)

}

func (log myLogger) Error(format string, a ...interface{}) {
	logPri(log.logLevel, ERROR, format, a)

}

func (log myLogger) Fatal(format string, a ...interface{}) {

	logPri(log.logLevel, FATAL, format, a)

}

func timeParse() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}
func logPri(logLevel, level logLevel, format string, a ...interface{}) {
	if judge(logLevel, level) {

		sli := info(3)
		msg := fmt.Sprintf(format, a...)
		line, _ := strconv.Atoi(sli[2])
		fmt.Printf("[%s] [%s] [%s:%s:%d] [%s]\n", timeParse(),
			unParse(level), sli[1], sli[0], line, msg)
	}
}
func info(skip int) []string {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("get info failed")
		return nil
	}

	methodName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[0]
	basePath := path.Base(file)

	sli := make([]string, 0, 3)
	sli = append(append(append(sli, methodName), basePath), strconv.Itoa(line))
	return sli
}

func unParse(level logLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case TRACE:
		return "TRACE"
	case WARRING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}

}
