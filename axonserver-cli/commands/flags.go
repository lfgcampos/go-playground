package commands

import (
	"github.com/urfave/cli/v2"
	"sort"
)

func Flags(app *cli.App) {
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "token",
			Usage: "Authentication Token",
			//Required:    true,
			Aliases:     []string{"t"},
			FilePath:    "token",
			EnvVars:     []string{"AXON_TOKEN"},
			Destination: &Token,
		},
		&cli.StringFlag{
			Name:        "server",
			Usage:       "URL of AxonServer",
			Aliases:     []string{"S"},
			Value:       Server,
			Destination: &Server,
		},
	}
	// sort Flags alphabetically
	sort.Sort(cli.FlagsByName(app.Flags))
}
