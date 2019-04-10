package main

import (
	"log"
	"os"

	"github.com/ispec-inc/micro-gen/cmd"
	"github.com/micro/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "Language for the greeting",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "complete a task on the list",
			Action:  cmd.Init,
		},
		{
			Name:    "scaffold",
			Aliases: []string{"s"},
			Usage:   "add a task to the list",
			Action:  cmd.Scaffold,
		},
		{
			Name:    "entity",
			Aliases: []string{"e"},
			Usage:   "add a task to the list",
			Action:  cmd.Entity,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
