package commands

import (
	"strings"

	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/io/log"
)

type dockerCommand struct {
	executor     Executor
	interpolator io.Interpolator
	ctx          config.Context
}

func (d dockerCommand) Build(tagName string) error {
	_, err := d.executeCommand("Build", "build", "--no-cache", "--tag", tagName, ".")
	return err
}

func (d dockerCommand) executeCommand(logName string, args ...string) (output []byte, err error) {
	signal := log.StartTrace(logName)
	defer log.StopTrace(signal, err)

	interpolationValues := map[string]string{
		"PWD": d.ctx.TargetPath,
	}

	command := d.interpolator.SubstituteValues(interpolationValues, d.ctx.BuildCommand)

	cmd := strings.Fields(command)
	if len(cmd) > 1 {
		args = append(cmd[1:], args...)
		command = cmd[0]
	}

	log.Debugf("Executing: %s %v", command, args)
	output, err = d.executor.Run(command, args...)
	return
}
