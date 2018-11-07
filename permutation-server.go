package main

import (
	"log"
	"os"

	"github.com/urfave/cli"

	"permutation-server/cmd"
	"permutation-server/pkg/settings"
)

func initApp(c *cli.Context) error {
	if c.IsSet("verbose") {
		settings.VerboseMode = true
	}

	defaultCfg := "/etc/rpc.conf"
	if _, err := os.Stat(defaultCfg); err == nil {
		err := settings.FromFile(defaultCfg)
		if err != nil {
			log.Fatal("initApp: get settings from file error.")
		}
	}

	if c.IsSet("config") {
		if len(c.String("config")) != 0 {
			err := settings.FromFile(c.String("config"))
			if err != nil {
				log.Fatal("initApp: get settings from file error.")
			}
		}
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "permutation-server"
	app.Usage = "Permutation server"
	app.Commands = []cli.Command{
		cmd.Srv,
		cmd.Adm,
	}
	app.Flags = append(app.Flags, cli.StringFlag{
		Name: "config, c", Value: "/etc/permutation.conf",
		Usage: "Load configuration `FILE`",
	})
	app.Flags = append(app.Flags, cli.BoolFlag{
		Name:  "verbose, vv",
		Usage: "Enable verbose mode",
	})
	app.Before = initApp
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
