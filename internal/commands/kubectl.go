package commands

import (
	"fmt"

	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io/log"
)

type kubectlCommand struct {
	executor Executor
}

func (k kubectlCommand) Deploy(env string, manifestPath string, ctx config.Context) (err error) {
	cmd := fmt.Sprintf("%s apply -f %s", ctx.DeployCommand, manifestPath)
	return k.executeCommand(cmd, "Deploy")
}

func (k kubectlCommand) Diff(env string, manifestPath string, ctx config.Context) (err error) {
	cmd := fmt.Sprintf("%s diff -f %s", ctx.DeployCommand, manifestPath)
	return k.executeCommand(cmd, "Diff")
}

func (k kubectlCommand) Show(env string, ctx config.Context) error {
	return nil
}

func (k kubectlCommand) executeCommand(command string, logName string) (err error) {
	signal := log.StartTrace(logName)
	defer log.StopTrace(signal, err)

	err = k.executor.Run(command)
	return err
}
