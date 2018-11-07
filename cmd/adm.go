package cmd

import (
	"github.com/urfave/cli"

	"permutation-server/pkg/database"
)

var Adm = cli.Command{
	Name:        "adm",
	Usage:       "Administrative tools for rpc server",
	Description: ``,
	Subcommands: []cli.Command{
		{
			Name:   "initdb",
			Usage:  "Initialize database",
			Action: initDB,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "force, f",
					Usage: "Think twice!"},
			},
		},
	},
}

func initDB(c *cli.Context) error {
	db, err := database.NewDB()
	defer db.Close()
	if err != nil {
		return err
	}
	return db.Init(c.Bool("force"))
}
