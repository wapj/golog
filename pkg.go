package golog

func Debug(v ...interface{}) {
	Logger.Log(DEBUG, v...)
}

func Info(v ...interface{}) {
	Logger.Log(INFO, v...)
}

func Warn(v ...interface{}) {
	Logger.Log(WARN, v...)
}

func Error(v ...interface{}) {
	Logger.Log(ERROR, v...)
}

func Fatal(v ...interface{}) {
	Logger.Log(FATAL, v...)
}

func Debugf(msg string, v ...interface{}) {
	Logger.Logf(DEBUG, msg, v...)
}

func Infof(msg string, v ...interface{}) {
	Logger.Logf(INFO, msg, v...)
}

func Warnf(msg string, v ...interface{}) {
	Logger.Logf(WARN, msg, v...)
}

func Errorf(msg string, v ...interface{}) {
	Logger.Logf(ERROR, msg, v...)
}

func Fatalf(msg string, v ...interface{}) {
	Logger.Logf(FATAL, msg, v...)
}