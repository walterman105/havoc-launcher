package util

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/process"
)

func InjectDll(processId int, dllPath string) {
	injectorPath := WorkingDir + "\\Injector.exe"
	pid := strconv.Itoa(int(processId))

	cmd := exec.Command(injectorPath, "--process-id", pid, "--inject", dllPath)

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Injected Successfully")
}

func DelayInjectDll(processID int, dllPath string, delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
	InjectDll(processID, dllPath)
}

func SuspendProcess(pid int) error {
	process, err := process.NewProcess(int32(pid))
	if err != nil {
		return err
	}

	err = process.Suspend()
	if err != nil {
		return err
	}

	fmt.Printf("Process with PID %d suspended.\n", pid)
	return nil
}

func KillProcessByName(processName string) error {
	cmd := exec.Command("taskkill", "/F", "/IM", processName)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to kill process '%s': %w", processName, err)
	}
	return nil
}
