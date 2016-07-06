package golog

import (
	"os"
	"log"
	"runtime"
	"path/filepath"
	"time"
	"strconv"
	"fmt"
	"strings"
)

/*
golog is easy to use go logger.
if you want to find simple and easy go log alternative
then this is what you find.
 */

type logger struct {
	level int
	l *log.Logger
}

const (
	FATAL = 50 - 10 * iota
	ERROR
	WARN
	INFO
	DEBUG
)

var (
	Logger *logger = &logger{ERROR, log.New(os.Stderr, "", 0)}
)

type CallerInfo struct {
	FileName string
	FilePath string
	FuncName string
	LineNo	int
}

func (l *logger) SetLogger(level string, pLogger *log.Logger ){
	var lvl int = getLevelFromName(strings.ToUpper(level))
	Logger = &logger{lvl, pLogger}
}

func getCallerInfo(depth int) *CallerInfo {
	if pc, filePath, lineNo, ok := runtime.Caller(depth + 1); !ok {
		return &CallerInfo{"unknown_file", "unknown_path", "unknown_func", 0}
	} else {
		return &CallerInfo{filepath.Base(filePath), filePath, runtime.FuncForPC(pc).Name(), lineNo}
	}
}

func getLevelName(level int) string {
	switch level {
	case FATAL:
		return "FATAL"
	case ERROR:
		return "ERROR"
	case WARN:
		return "WARN"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	default:
		return "ERROR"
	}
}

func getLevelFromName(level string) int {
	switch level {
	case "FATAL":
		return FATAL
	case "ERROR":
		return ERROR
	case "WARN":
		return WARN
	case "INFO":
		return INFO
	case "DEBUG":
		return DEBUG
	default:
		return ERROR
	}
}

func getPrefix(level string) string {
	caller := getCallerInfo(3)

	s := []string{"[" + level + "]",
		time.Now().Format("2006-01-02 15:04:05"),
		caller.FileName + ":" + strconv.Itoa(caller.LineNo),
		caller.FuncName}

	return strings.Join(s, " ")

}

func (l *logger)Log(level int, v ...interface{}){
	if l.level <= level {
		l.l.Println(getPrefix(getLevelName(level)),	fmt.Sprint(v...))
	}
}

func (l *logger)Logf(level int, msg string, v ...interface{}) {
	if l.level <= level {
		l.l.Print(getPrefix(getLevelName(level)), " ", fmt.Sprintf(msg, v...))
	}
}

func (l *logger)Debug(v ...interface{}) {
	l.Log(DEBUG, v...)
}

func (l *logger)Info(v ...interface{}) {
	l.Log(INFO, v...)
}

func (l *logger)Warn(v ...interface{}) {
	l.Log(WARN, v...)
}

func (l *logger)Error(v ...interface{}) {
	l.Log(ERROR, v...)
}

func (l *logger)Fatal(v ...interface{}) {
	l.Log(FATAL, v...)
}

func (l *logger)Debugf(msg string, v ...interface{}) {
	l.Logf(DEBUG, msg, v...)
}

func (l *logger)Infof(msg string, v ...interface{}) {
	l.Logf(INFO, msg, v...)
}

func (l *logger)Warnf(msg string, v ...interface{}) {
	l.Logf(WARN, msg, v...)
}

func (l *logger)Errorf(msg string, v ...interface{}) {
	l.Logf(ERROR, msg, v...)
}

func (l *logger)Fatalf(msg string, v ...interface{}) {
	l.Logf(FATAL, msg, v...)
}

