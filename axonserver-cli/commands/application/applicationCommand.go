package application

import (
	"github.com/urfave/cli/v2"
)

//
func ApplicationCommand(app *cli.App) {
	var userCommand = []*cli.Command{
		{
			Name:        "application",
			Aliases:     []string{"a"},
			Usage:       "commands related to applications",
			Description: "This is the command related to applications",
		},
	}
	app.Commands = append(app.Commands, userCommand...)
}
