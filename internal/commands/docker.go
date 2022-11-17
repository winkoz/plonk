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

func (d dockerCommand) Build(tagName string, skipCache bool) error {
	// We build with a unique tag and if specified the latest
	tags := []string{
		"--tag",
		tagName,
	}

	// Docker command line arguments
	args := []string{"buildx", "build"}

	if !skipCache {
		args = append(args, "--no-cache")
	}

	args = append(args, tags...)
	args = append(args, "--platform", "linux/arm64,linux/amd64,linux/amd64/v2")
	args = append(args, "--push")
	args = append(args, ".")

	_, err := d.executeCommand("Build", args...)
	return err
}

func (d dockerCommand) Push(tagName string) error {
	// Docker command line arguments
	args := []string{
		"push",
		tagName,
	}
	_, err := d.executeCommand("Push", args...)
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
