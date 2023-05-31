package main

import (
	"havoc-launcher/gui"
	"havoc-launcher/launcher"
	"havoc-launcher/util"
)

func main() {
	util.SetWorkingDir()
	util.ReadConfig()
	launcher.UpdateLaunchArgs()
	gui.StartGUI()
}
