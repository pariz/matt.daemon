package process

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pariz/matt.daemon/config"
)

var cfg config.Config

func Init(c config.Config) {
	cfg = c
	for processName := range cfg.Processes {

		process := cfg.Processes[processName]

		// Check if process exists
		if processExists(processName) == false || processRunning(processName) == false {
			go spawnProcess(processName, process)
			continue
		}

	}

}

func processExists(processName string) bool {
	return false
}

func processRunning(processName string) bool {
	return false
}

func createPid(processName string, pid int) (err error) {
	var file *os.File

	file, err = os.Create(fmt.Sprintf("%s/%s.pid", cfg.PidDir, processName))

	if err != nil {
		return
	}

	file.WriteString(fmt.Sprintf("%d", pid))
	file.Close()

	return
}

func createStdLogFile(processName, stdType string) (file *os.File) {

	filePath := fmt.Sprintf("%s/%s.%s.log", cfg.LogDir, processName, stdType)
	var err error

	if _, err = os.Stat(filePath); err != nil {
		fmt.Println("Create log file", filePath)
		file, err = os.OpenFile(filePath, os.O_CREATE, 0755)
		fmt.Printf("error creating logfile: %s\n", err)
	} else {
		fmt.Println("Using existing logfile", filePath)
		file, err = os.OpenFile(filePath, os.O_RDWR, 0755)
		fmt.Printf("error opening logfile: %s\n", err)
	}

	return file

}

func spawnProcess(processName string, process config.Process) {
	fmt.Printf("Spawning process %s\n", processName)

	done := make(chan error, 1)

	var (
		cmd *exec.Cmd
	)

	cmd = exec.Command(process.Script)

	// Log stdout
	cmd.Stdout = createStdLogFile(processName, "stdout")
	cmd.Stderr = createStdLogFile(processName, "stderr")

	err := cmd.Start()

	if err = createPid(processName, cmd.Process.Pid); err != nil {
		fmt.Printf("Could not create pid, %s, %d. Err: %s", processName, cmd.Process.Pid, err)
	}

	// wait for exit
	done <- cmd.Wait()

	select {
	case err = <-done:
		if err != nil {
			// Spawn the process again. Todo: Check for retries
			fmt.Printf("Process %s died.. restarting\n", processName)

			go spawnProcess(processName, process)

		} else {
			fmt.Printf("Process %s terminated gracefully\n", processName)
		}
	}

	if err != nil {
		return
	}

}
