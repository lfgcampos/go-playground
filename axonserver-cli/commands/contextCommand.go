package commands

import "github.com/urfave/cli/v2"

func ContextCommand(app *cli.App) {
	var userCommand = []*cli.Command{
		{
			Name:        "context",
			Aliases:     []string{"c"},
			Usage:       "commands related to contexts",
			Description: "This is the command related to contexts",
		},
	}
	app.Commands = append(app.Commands, userCommand...)
}
