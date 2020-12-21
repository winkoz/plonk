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
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	outputBytes := cmdOutput.Bytes()
	if err != nil {
		runErr := fmt.Errorf("[Internal Error] Command could not be executed successfully. %s.\n\terror = %v", string(outputBytes), err)
		log.Errorf(runErr.Error())
		return nil, runErr
	}
	log.Infof("[INFO] Executed command:\n\t%s", string(outputBytes))
	return outputBytes, nil
}
