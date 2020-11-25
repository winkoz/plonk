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

func (k kubectlCommand) Deploy(env string, manifestPath string) (err error) {
	return k.executeCommand("Deploy", "apply", "-f", manifestPath)
}

func (k kubectlCommand) Diff(env string, manifestPath string) (err error) {
	return k.executeCommand("Diff", "diff", "-f", manifestPath)
}

func (k kubectlCommand) Show(env string) error {
	return nil
}

func (k kubectlCommand) executeCommand(logName string, args ...string) (err error) {
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
	err = k.executor.Run(command, args...)
	return err
}
