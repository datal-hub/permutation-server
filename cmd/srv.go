package cmd

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/cli"

	"permutation-server/handlers"
	"permutation-server/pkg/settings"
)

var Srv = cli.Command{
	Name:        "srv",
	Usage:       "Start permutation server",
	Description: ``,
	Action:      runSrv,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "listen, l", Value: "127.0.0.1:8181",
			Usage: "Start server at specified address, port",
		},
	},
}

func updateSettings(c *cli.Context) {
	if c.IsSet("listen") {
		settings.ListenAddr = c.String("listen")
	}
}

func runSrv(c *cli.Context) error {
	updateSettings(c)

	r := mux.NewRouter()

	r.Methods("POST").Path("/v1/init").Handler(http.HandlerFunc(handlers.Init))
	r.Methods("GET").Path("/v1/next").Handler(http.HandlerFunc(handlers.Next))

	err := http.ListenAndServe(settings.ListenAddr, r)
	log.Fatal(err)

	return nil
}
