package commands

import (
	"fmt"

	"github.com/winkoz/plonk/internal/config"
)

type kubecltCommand struct {
	executor Executor
}

func (k kubecltCommand) Deploy(env string, manifestPath string, ctx config.Context) error {
	cmd := fmt.Sprintf("kubectl -f %s", manifestPath)
	err := k.executor.Run(cmd)
	return err
}

func (k kubecltCommand) Show(env string, ctx config.Context) error {
	return nil
}
