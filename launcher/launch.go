package launcher

import (
	"fmt"
	"havoc-launcher/util"
	"log"
	"os/exec"
)

func StartFortnite(useFnLauncher bool) {
	//Old versions of Fortnite (before ~Season 3) don't have a FortniteLauncher.exe
	var fnlauncher *exec.Cmd

	if useFnLauncher {
		fmt.Println("Starting Fortnite Launcher...")
		fnlauncher = exec.Command(FortnitePath + LauncherPath)
		err := fnlauncher.Start()
		if err != nil {
			log.Fatal(err)
		}

		err = util.SuspendProcess(fnlauncher.Process.Pid)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Skipping Fortnite Launcher...")
	}

	fmt.Println("Starting Battle Eye...")
	anticheat := exec.Command(FortnitePath + CheatPath)
	err := anticheat.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = util.SuspendProcess(anticheat.Process.Pid)
	if err != nil {
		log.Fatal(err)
	}

	UpdateLaunchArgs()

	fmt.Println("Starting Fortnite...")
	fortnite := exec.Command(FortnitePath+ClientPath, LaunchArgs...)
	err = fortnite.Start()
	if err != nil {
		log.Fatal(err)
	}

	//Kill Battle Eye and the FortniteLauncher after the game closes
	go func() {
		err := fortnite.Wait()
		if err != nil {
			fmt.Println("Error waiting for Fortnite to complete:", err)
		}

		fmt.Println("Fortnite has closed")
		if useFnLauncher {
			fnlauncher.Process.Kill()
		}
		anticheat.Process.Kill()
	}()

	fmt.Println("Injecting Cobalt")
	util.InjectDll(fortnite.Process.Pid, util.WorkingDir+"\\dll\\cobalt.dll")

	//TODO: Fix Injecting
	if util.GetConfig().Options.Console {
		fmt.Println("Injecting Console")
		go func() {
			util.DelayInjectDll(fortnite.Process.Pid, util.WorkingDir+"\\dll\\consolev2.dll", 25)
		}()
	}

	if util.GetConfig().Options.MemoryFix {
		fmt.Println("Injecting Memory Fix")
		go func() {
			util.DelayInjectDll(fortnite.Process.Pid, util.WorkingDir+"\\dll\\leakv2.dll", 30)
		}()
	}
}
