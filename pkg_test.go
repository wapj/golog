package golog

import "testing"

func TestPkg(t *testing.T) {
	Debug("DEBUG")
	Info("Info")
	Warn("Warn")
	Error("Error")
	Fatal("Fatal")

	SetLevel("debug")

	Debugf("%d test", 10)
	Infof("%d info test", 10)
	Warnf("%d warning", 10)
	Errorf("%d error", 10)
	Fatalf("%d fatal", 10)
}