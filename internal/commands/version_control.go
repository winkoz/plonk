package commands

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/winkoz/plonk/internal/config"
)

type gitVersionControllerCommand struct {
	ctx config.Context
}

func (g gitVersionControllerCommand) Head() (string, error) {
	fmt.Printf("%v+", g.ctx.TargetPath)
	r, err := git.PlainOpen(g.ctx.TargetPath)
	if err != nil {
		return "", err
	}
	head, err := r.Head()
	if err != nil {
		return "", err
	}
	return head.Hash().String(), nil
}
