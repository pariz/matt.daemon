package process

import "errors"

type Rpc struct{}

func (r *Rpc) Kill(processName *string, reply *string) error {

	var (
		err error
	)

	err = killProcess(*processName)

	if err != nil {
		return err
	}

	*reply = "The process was killed"

	return nil

}

func (r *Rpc) Start(processName *string, reply *string) error {

	if pid, _ := getPidForProcess(*processName); pidRunning(pid) == true {
		return errors.New("The process is already running")
	}

	if processCfg, ok := cfg.Processes[*processName]; ok {
		spawnProcess(*processName, processCfg)
		*reply = "The process was started"
	} else {
		return errors.New("The process does not exists in the process config")
	}

	return nil

}
