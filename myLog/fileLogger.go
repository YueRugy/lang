package myLog

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	filePath  string
	fileName  string
	maxSize   uint64
	level     logLevel
	fh, errFh *os.File
}

var logChan chan *logMsg

type logMsg struct {
	basePath   string
	methodName string
	msg        string
	line       int
	fileLogger *FileLogger
	level      logLevel
}

func init() {
	logChan = make(chan *logMsg, 100000)
}

//构造函数
func NewFileLogger(fp, fn, level string, ms uint64) *FileLogger {
	logLevel, err := parseLevel(level)
	if err != nil {
		return nil
	}
	fh, fhErr := initFile(fp, fn)
	return &FileLogger{
		filePath: fp,
		fileName: fn,
		maxSize:  ms,
		level:    logLevel,
		fh:       fh,
		errFh:    fhErr,
	}

}

func initFile(fp, fn string) (*os.File, *os.File) {
	fullName := path.Join(fp, fn)
	fh, err := os.OpenFile(fullName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0643)
	if err != nil {
		fmt.Printf("open file failed %v", err)
		recover()
	}
	errName := fullName + ".err"
	fhErr, err := os.OpenFile(errName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0643)
	if err != nil {
		fmt.Printf("open err file failed %v", err)
		recover()
	}
	//这里开始写日志
	for i := 0; i < 5; i++ {
		go writeLogBack()
	}
	return fh, fhErr
}

func (fl *FileLogger) checkSize(file *os.File) bool {
	info, err := file.Stat()
	if err != nil {
		fmt.Printf("%v", err)
	}
	return uint64(info.Size()) >= fl.maxSize
}

func (fl *FileLogger) spiltFile(file *os.File) (*os.File, error) {

	var newFile *os.File

	str := time.Now().Format("2006010215040500")
	fName := file.Name()
	file.Close()
	err := os.Rename(fName, fName+str)
	/*if err != nil {
		fmt.Printf("rename file failed %v", err)
		return file, err
	}*/
	file.Close()

	newFile, err = os.OpenFile(fName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new file failed %v", err)
		return nil, err
	}
	return newFile, nil
}
func fileLog(fl *FileLogger, level logLevel, format string, a ...interface{}) {
	if judge(fl.level, level) {
		methodName, basePath, line := info(3)
		msg := fmt.Sprintf(format, a...)

		logMsg := &logMsg{
			basePath:   basePath,
			methodName: methodName,
			msg:        msg,
			line:       line,
			fileLogger: fl,
			level:      level,
		}
		//通道满的时候丢失日志
		select {
		case logChan <- logMsg:
		default:
			fmt.Println("丢失日志")
		}

		//line, _ := strconv.Atoi(sli[2])

	}
}
func writeLogBack() {
	for {
		select {
		case log := <-logChan:
			if log.fileLogger.checkSize(log.fileLogger.fh) {
				newFile, err := log.fileLogger.spiltFile(log.fileLogger.fh)
				if err != nil {
					fmt.Printf("%v", err)
					return
				}
				log.fileLogger.fh = newFile
			}
			_, err := fmt.Fprint(log.fileLogger.fh, fmt.Sprintf("[%s] [%s] [%s:%s:%d] [%s]\n", timeParse(),
				unParse(log.level), log.basePath, log.methodName, log.line, log.msg))
			if err != nil {
				fmt.Printf("write log to file failed%v", err)
				return
			}

			if log.level >= ERROR {
				if log.fileLogger.checkSize(log.fileLogger.errFh) {
					newFile, err := log.fileLogger.spiltFile(log.fileLogger.errFh)
					if err != nil {
						fmt.Printf("%v", err)
						return
					}
					log.fileLogger.errFh = newFile
				}
				_, err = fmt.Fprint(log.fileLogger.errFh, fmt.Sprintf("[%s] [%s] [%s:%s:%d] [%s]\n", timeParse(),
					unParse(log.level), log.basePath, log.methodName, log.line, log.msg))
				if err != nil {
					fmt.Printf("write errlog to file failed%v", err)
					return
				}
			}
		default:
			//没有日志要写是让出cpu资源
			time.Sleep(time.Millisecond * 500)

		}
	}

}

func (fl *FileLogger) Debug(format string, a ...interface{}) {
	fileLog(fl, DEBUG, format, a)
}

func (fl *FileLogger) Trace(format string, a ...interface{}) {
	fileLog(fl, TRACE, format, a)
}

func (fl *FileLogger) Info(format string, a ...interface{}) {
	fileLog(fl, INFO, format, a)
}
func (fl *FileLogger) Warning(format string, a ...interface{}) {
	fileLog(fl, WARRING, format, a)
}

func (fl *FileLogger) Error(format string, a ...interface{}) {
	fileLog(fl, ERROR, format, a)
}
func (fl *FileLogger) Fatal(format string, a ...interface{}) {
	fileLog(fl, FATAL, format, a)
}

func (fl *FileLogger) FClose() {
	fl.errFh.Close()
	fl.fh.Close()
}
