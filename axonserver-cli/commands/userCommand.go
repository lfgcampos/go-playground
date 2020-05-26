package commands

import "github.com/urfave/cli/v2"

func UserCommand(app *cli.App) {
	var userCommand = []*cli.Command{
		{
			Name:        "user",
			Aliases:     []string{"u"},
			Usage:       "commands related to users",
			Description: "This is the command related to users",
		},
	}
	app.Commands = append(app.Commands, userCommand...)
}
