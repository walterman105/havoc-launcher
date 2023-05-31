package launcher

import (
	"fmt"
	"havoc-launcher/util"
	"os"
)

var FortnitePath string

var LauncherPath = "\\FortniteGame\\Binaries\\Win64\\FortniteLauncher.exe"
var CheatPath = "\\FortniteGame\\Binaries\\Win64\\FortniteClient-Win64-Shipping_BE.exe"
var ClientPath = "\\FortniteGame\\Binaries\\Win64\\FortniteClient-Win64-Shipping.exe"

var LaunchArgs []string

var Username string
var Password string

func UpdateLaunchArgs() {
	Username, Password = util.GetCredentials()

	LaunchArgs = []string{
		"-epicapp=Fortnite",
		"-epicenv=Prod",
		"-epiclocale=en-us",
		"-epicportal",
		"-noeac",
		"-nobe",
		"-skippatchcheck",
		"-fltoken=e8eb05fag41046i3hd23c89c",
		"-AUTH_TYPE=epic",
		"-AUTH_LOGIN=" + Username + "@Havoc.Launcher",
		"-AUTH_PASSWORD=" + Password,
	}
}

func fileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("File '%s' exists.\n", filePath)
		return true
	} else if os.IsNotExist(err) {
		fmt.Printf("File '%s' does not exist.\n", filePath)
		return false
	} else {
		fmt.Printf("Error checking file: %v\n", err)
		return false
	}
}

func VerifyFnPath(fnPath string) bool {
	fortniteClient := fnPath + ClientPath

	if fileExists(fortniteClient) {
		FortnitePath = fnPath
		return true
	} else {
		return false
	}
}

func GetLauncherType() bool {
	return fileExists(FortnitePath + LauncherPath)
}
