package main

import (
	"os"

	lcli "github.com/aliensero/go-lotus-interaction/cli"
	"github.com/prometheus/common/log"
	"github.com/urfave/cli/v2"
)

func main() {

	local := []*cli.Command{
		ClientDealCmd,
	}

	app := &cli.App{
		Name:    "lotus-interaction",
		Usage:   "lotus interaction test",
		Version: "0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Hidden:  true,
				Value:   "~/.lotus",
			},
		},
		Commands: local,
	}

	if err := app.Run(os.Args); err != nil {
		_, ok := err.(*lcli.ErrCmdFailed)
		if ok {
			log.Debugf("%+v", err)
		} else {
			log.Warnf("%+v", err)
		}
		os.Exit(1)
	}

}
