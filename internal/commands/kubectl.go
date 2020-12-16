package commands

import (
	"strings"

	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

type kubectlCommand struct {
	executor     Executor
	interpolator io.Interpolator
	ctx          config.Context
}

func (k kubectlCommand) Deploy(manifestPath string) error {
	_, err := k.executeCommand("Deploy", "apply", "-f", manifestPath)
	return err
}

func (k kubectlCommand) Diff(manifestPath string) error {
	_, err := k.executeCommand("Diff", "diff", "-f", manifestPath)
	return err
}

func (k kubectlCommand) Show(env string) error {
	return nil
}

func (k kubectlCommand) GetPods(env string) ([]byte, error) {
	return nil, nil
}

func (k kubectlCommand) executeCommand(logName string, args ...string) (output []byte, err error) {
	signal := log.StartTrace(logName)
	defer log.StopTrace(signal, err)

	interpolationValues := map[string]string{
		"PWD": k.ctx.TargetPath,
	}

	command := k.interpolator.SubstituteValues(interpolationValues, k.ctx.DeployCommand)

	cmd := strings.Fields(command)
	if len(cmd) > 1 {
		args = append(cmd[1:], args...)
		command = cmd[0]
	}

	log.Debugf("Executing: %s %v", command, args)
	output, err = k.executor.Run(command, args...)
	return
}
