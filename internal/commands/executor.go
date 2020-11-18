package commands

import (
	"bytes"
	"os/exec"

	"github.com/winkoz/plonk/internal/io/log"
)

// Executor command line executioner
type Executor interface {
	Run(command string) error
}

// NewExecutor factory for command line executors
func NewExecutor() Executor {
	return executor{}
}

type executor struct {
}

func (e executor) Run(command string) error {
	cmd := exec.Command("echo", command)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		log.Errorf("Command could not be executed successfully. error = %v", err)
		return err
	}
	log.Infof("Executed command: %s", string(cmdOutput.Bytes()))
	return nil
}
