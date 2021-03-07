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

func (d dockerCommand) Build(tagName string, isLatest bool) error {
	// We build with a unique tag and if specified the latest
	tags := []string{
		"--tag",
		tagName,
	}
	if isLatest {
		tags = append(tags, "--tag")
		tags = append(tags, "latest")
	}

	// Docker command line arguments
	args := []string{
		"build",
		"--no-cache",
	}
	args = append(args, tags...)
	args = append(args, ".")

	_, err := d.executeCommand("Build", args...)
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
