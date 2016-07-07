package golog

import (
	"testing"
	"os"
	"io"
	"log"
)


func getTesLogger() *log.Logger{
	logFile, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Error Can't open file ", err)
	}

	fileAndStdWriter := io.MultiWriter(logFile, os.Stdout)
	return log.New(fileAndStdWriter, "", 0)
}

func TestLogger_Debug(t *testing.T) {
	// doesn't print cause default logging level is ERROR
	Debug("TEST Debug")
	Debugf("%d test", 10)

	// initialize Logger
	SetLogger("debug", getTesLogger())

	// and test again it does work
	Debug("TEST Debug")
	Debugf("%d test", 10)

	SetLogger("warn", getTesLogger())
}


