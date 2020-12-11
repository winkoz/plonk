package commands

import (
	"bytes"
	"fmt"
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
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		runErr := fmt.Errorf("[Internal Error] Command could not be executed successfully. %s.\n\terror = %v.", string(cmdOutput.Bytes()), err)
		log.Errorf(runErr.Error())
		return runErr
	}
	log.Infof("[INFO] Executed command:\n\t%s", string(cmdOutput.Bytes()))
	return nil
}
