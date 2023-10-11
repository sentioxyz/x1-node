package main

import (
	"github.com/okx/zkevm-node/test/scripts/cmd/compilesc"
	"github.com/urfave/cli/v2"
)

func compileSC(ctx *cli.Context) error {
	manager, err := compilesc.NewManager(ctx.String(flagInput))
	if err != nil {
		return err
	}

	return manager.Run()
}
