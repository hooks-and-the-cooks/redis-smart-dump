package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func createCLIApp(logger *logrus.Logger) cli.App {
	app := cli.App{
		Name:    "redis-smart-dump",
		Usage:   "Dump and restoreKeys keys in Redis Cache",
		Version: "0.1.0",
		Commands: []*cli.Command{
			{
				Name:    "dump",
				Aliases: []string{"d"},
				Usage:   "Dump Keys from Redis",
				Action: func(ctx *cli.Context) error {
					host := ctx.Args().Get(0)
					port := ctx.Args().Get(1)
					dumpKeys(host, port, logger)
					return nil
				},
			},
			{
				Name:    "restore",
				Aliases: []string{"r"},
				Usage:   "restore keys to Redis",
				Action: func(ctx *cli.Context) error {
					host := ctx.Args().Get(0)
					port := ctx.Args().Get(1)
					fileName := ctx.Args().Get(2)
					restoreKeys(fileName, host, port, logger)
					return nil
				},
			},
		},
	}
	return app
}
