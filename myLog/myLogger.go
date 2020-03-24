package myLog

import (
	"errors"
	"fmt"
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
func judge(log myLogger, level logLevel) bool {
	return log.logLevel <= level
}
func (log myLogger) Debug(msg string) {
	if judge(log, DEBUG) {
		fmt.Printf("[%s] [DEBUG] %s\n", timeParse(), msg)
	}
}

func (log myLogger) Trace(msg string) {
	if judge(log, TRACE) {
		fmt.Printf("[%s] [DEBUG] %s\n", timeParse(), msg)
	}
}

func (log myLogger) Info(msg string) {
	if judge(log, INFO) {
		fmt.Printf("[%s] [DEBUG] %s\n", timeParse(), msg)
	}
}

func (log myLogger) Warning(msg string) {
	if judge(log, WARRING) {
		fmt.Printf("[%s] [DEBUG] %s\n", timeParse(), msg)
	}
}

func (log myLogger) Error(msg string) {
	if judge(log, ERROR) {
		fmt.Printf("[%s] [DEBUG] %s\n", timeParse(), msg)
	}
}

func (log myLogger) Fatal(msg string) {
	if judge(log, FATAL) {
		fmt.Printf("[%s] [DEBUG] %s\n", timeParse(), msg)
	}
}

func timeParse() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}
