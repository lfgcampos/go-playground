package main

import (
	"github.com/urfave/cli/v2"
	"io.axoniq.cli/commands"
	"io.axoniq.cli/info"
	"log"
	"os"
)

var app = cli.NewApp()

func main() {
	info.AppInfo(app)
	commands.Flags(app)

	// users
	commands.UserCommand(app)
	commands.UserListSubCommand(app)
	commands.UserRegisterSubCommand(app)

	// contexts
	commands.ContextCommand(app)
	commands.ContextListSubCommand(app)

	// applications
	commands.ApplicationCommand(app)
	commands.ApplicationListSubCommand(app)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
