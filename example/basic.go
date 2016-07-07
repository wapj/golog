package main

import "github.com/wapj/golog"

func main() {

	// not printed below cause default logging level is warning
	golog.Info("Also will not print")
	golog.Warn("Warning! warning! this is warning message!")
	golog.Error("Use it carefully for find error")
	golog.Fatal("This is emergency situation! You have to resolve this problem right now!")

}