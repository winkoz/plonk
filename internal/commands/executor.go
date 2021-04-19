package commands

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

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
	var stdoutBuf, stderrBuf bytes.Buffer

	cmd := exec.Command(command, args...)
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	log.Infof("Executing command: %s %s", command, strings.Join(args, " "))
	err := cmd.Run()
	if err != nil {
		errOutput := stderrBuf.String()
		runErr := fmt.Errorf("command could not be executed successfully. error = %v\n\t%s", err, errOutput)
		log.Errorf(runErr.Error())
		return nil, runErr
	}

	cmdOutput := stdoutBuf.String()
	log.Infof("Command executed successfully. output = \n%s\n", cmdOutput)

	log.Infof("Executed command: %s %s", command, strings.Join(args, " "))
	return []byte(cmdOutput), nil
}
