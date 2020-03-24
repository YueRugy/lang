package myLog

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

type fileLogger struct {
	filePath  string
	fileName  string
	maxSize   uint64
	level     logLevel
	fh, errFh *os.File
}

//构造函数
func NewFileLogger(fp, fn, level string, ms uint64) *fileLogger {
	logLevel, err := parseLevel(level)
	if err != nil {
		return nil
	}
	fullName := path.Join(fp, fn)
	fh, err := os.OpenFile(fullName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed %v", err)
		return nil
	}
	errName := fullName + ".err"
	fhErr, err := os.OpenFile(errName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err file failed %v", err)
		return nil
	}

	return &fileLogger{
		filePath: fp,
		fileName: fn,
		maxSize:  ms,
		level:    logLevel,
		fh:       fh,
		errFh:    fhErr,
	}
}

func fileLog(fl fileLogger, level logLevel, format string, a ...interface{}) {
	if judge(fl.level, level) {
		sli := info(3)
		msg := fmt.Sprintf(format, a...)
		line, _ := strconv.Atoi(sli[2])
		_, err := fmt.Fprint(fl.fh, fmt.Sprintf("[%s] [%s] [%s:%s:%d] [%s]\n", timeParse(),
			unParse(level), sli[1], sli[0], line, msg))
		if err != nil {
			fmt.Printf("write log to file failed%v", err)
			return
		}

		if level >= ERROR {

			_, err := fmt.Fprint(fl.errFh, fmt.Sprintf("[%s] [%s] [%s:%s:%d] [%s]\n", timeParse(),
				unParse(level), sli[1], sli[0], line, msg))

			if err != nil {
				fmt.Printf("write errlog to file failed%v", err)
				return
			}
		}

	}
}

func (fl fileLogger) Debug(format string, a ...interface{}) {
	fileLog(fl, DEBUG, format, a)
}

func (fl fileLogger) Trace(format string, a ...interface{}) {
	fileLog(fl, TRACE, format, a)
}

func (fl fileLogger) Info(format string, a ...interface{}) {
	fileLog(fl, INFO, format, a)
}
func (fl fileLogger) Warning(format string, a ...interface{}) {
	fileLog(fl, WARRING, format, a)
}

func (fl fileLogger) Error(format string, a ...interface{}) {
	fileLog(fl, ERROR, format, a)
}
func (fl fileLogger) Fatal(format string, a ...interface{}) {
	fileLog(fl, FATAL, format, a)
}

func (fl fileLogger) FClose() {
	fl.errFh.Close()
	fl.fh.Close()
}
