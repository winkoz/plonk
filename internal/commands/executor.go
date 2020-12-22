package commands

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/winkoz/plonk/internal/io/log"
)

// Executor command line executioner
type Executor interface {
	Run(command string, args ...string) ([]byte, error)
}

// NewExecutor factory for command line executors
func NewExecutor() Executor {
	return executor{}
}

type executor struct {
}

func (e executor) Run(command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	cmdOutput := &bytes.Buffer{}
	errOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Stderr = errOutput

	err := cmd.Run()
	if err != nil {
		errBytes := errOutput.Bytes()
		runErr := fmt.Errorf("Command could not be executed successfully. error = %v\n\t%s", err, string(errBytes))
		log.Errorf(runErr.Error())
		return nil, runErr
	}
	outputBytes := cmdOutput.Bytes()
	log.Infof("[INFO] Executed command:\n\t%s", string(outputBytes))
	return outputBytes, nil
}
