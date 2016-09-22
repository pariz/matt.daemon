package process

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/pariz/matt.daemon/config"
)

var cfg config.Config

func Init(c config.Config) {
	cfg = c
	for processName := range cfg.Processes {

		process := cfg.Processes[processName]

		// Check if process exists
		if _, err := getPidForProcess(processName); err == nil {
			go spawnProcess(processName, process)
			continue
		}

	}

}

func killProcess(processName string) error {

	process := cfg.Processes[processName]

	if process != nil {
		return process.Cmd.Process.Kill()
	}

	return errors.New("The process does not exist")

}

func getPidForProcess(processName string) (pid int, err error) {

	pidFilePath := fmt.Sprintf("%s/%s.pid", cfg.PidDir, processName)
	var pidBytes []byte

	if _, err = os.Stat(pidFilePath); os.IsNotExist(err) {

		return

	}

	pidBytes, err = ioutil.ReadFile(pidFilePath)

	if err != nil {
		return
	}

	pid, _ = strconv.Atoi(string(pidBytes))

	return
}

func killPid(pid int) error {
	return syscall.Kill(pid, 1)
}

func pidRunning(pid int) bool {

	_, err := os.FindProcess(pid)
	fmt.Println("pidRunning", pid, err)
	return err != nil

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

	if _, err = os.Stat(filePath); os.IsNotExist(err) {

		file, _ = os.Create(filePath)

	} else {

		file, _ = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	}

	return file

}

func spawnProcess(processName string, process *config.Process) {
	fmt.Printf("Spawning process.. %s\n", processName)

	done := make(chan error, 1)

	var (
		cmd *exec.Cmd
	)

	cmd = exec.Command(process.Script)

	process.Cmd = cmd

	stdoutFile := createStdLogFile(processName, "stdout")
	stdErrFile := createStdLogFile(processName, "stderr")

	// Log stdout
	cmd.Stdout = stdoutFile
	cmd.Stderr = stdErrFile

	err := cmd.Start()

	if err = createPid(processName, cmd.Process.Pid); err != nil {
		fmt.Printf("Could not create pid, %s, %d. Err: %s", processName, cmd.Process.Pid, err)
	}

	// wait for exit
	done <- cmd.Wait()

	select {
	case err = <-done:

		stdoutFile.Close()
		stdErrFile.Close()

		if err != nil {
			// Spawn the process again. Todo: Check for retries
			fmt.Printf("Process %s died.. restarting\n", processName)

			go spawnProcess(processName, process)

		} else {
			fmt.Printf("Process %s terminated gracefully!\n", processName)
		}
	}

	if err != nil {
		return
	}

}
