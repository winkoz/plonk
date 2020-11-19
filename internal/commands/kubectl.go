package commands

import (
	"strings"

	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io/log"
)

type kubectlCommand struct {
	executor Executor
}

func (k kubectlCommand) Deploy(env string, manifestPath string, ctx config.Context) (err error) {
	return k.executeCommand("Deploy", ctx.DeployCommand, "apply", "-f", manifestPath)
}

func (k kubectlCommand) Diff(env string, manifestPath string, ctx config.Context) (err error) {
	return k.executeCommand("Diff", ctx.DeployCommand, "diff", "-f", manifestPath)
}

func (k kubectlCommand) Show(env string, ctx config.Context) error {
	return nil
}

func (k kubectlCommand) executeCommand(logName string, command string, args ...string) (err error) {
	signal := log.StartTrace(logName)
	defer log.StopTrace(signal, err)

	cmd := strings.Fields(command)
	if len(cmd) > 1 {
		args = append(cmd[1:], args...)
		command = cmd[0]
	}

	log.Debugf("Executing: %s %v", command, args)
	err = k.executor.Run(command, args...)
	return err
}
