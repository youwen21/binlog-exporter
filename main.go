package main

import (
	"binlog_exporter/cmd/command"
	_ "binlog_exporter/conf"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "dev",
				Aliases: []string{"c"},
				Usage:   "develop test",
				Action:  command.Dev,
			},
			{
				Name:  "binlog-start",
				Usage: "start consume mysql binlog",

				Flags: []cli.Flag{
					&cli.StringFlag{Name: "host", Value: ""},
					&cli.IntFlag{Name: "port", Value: 3306},
					&cli.StringFlag{Name: "username", Value: ""},
					&cli.StringFlag{Name: "password", Value: ""},
					&cli.StringFlag{Name: "charset", Value: "utf8"},
					&cli.IntFlag{Name: "server_id", Value: 0},
				},

				Action: command.BinlogClient.StartBinlogClient,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
