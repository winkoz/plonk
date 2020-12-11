package commands

import (
	"os/exec"

	"github.com/winkoz/plonk/internal/io/log"
)

// Executor command line executioner
type Executor interface {
	Run(command string, args ...string) error
}

// NewExecutor factory for command line executors
func NewExecutor() Executor {
	return executor{}
}

type executor struct {
}

func (e executor) Run(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("[Internal Error] Command could not be executed successfully. %s, error = %v.", string(cmdOutput), err)
		return err
	}
	log.Debugf("[INFO] Executed command: %s", string(cmdOutput))
	return nil
}
