package main

import (
	"./core"
	"github.com/nurza/logo"
)

var (
	// Logging
	l *logo.Logger

	// Config
	configPath = "config.yaml"
)

func main() {
	// Logging
	println("Init logger...")
	l = core.InitLogger()
	l.Verbose("logger OK")

	// Config
	l.Verbose("Init config")
	core.InitConfig(configPath)
	l.Verbose("config OK")

	// Init Core
	core.Init()

}
