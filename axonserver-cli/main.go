package main

import (
	"github.com/urfave/cli/v2"
	"io.axoniq.cli/commands"
	"io.axoniq.cli/commands/application"
	"io.axoniq.cli/commands/context"
	"io.axoniq.cli/commands/user"
	"io.axoniq.cli/info"
	"log"
	"os"
)

var app = cli.NewApp()

func main() {
	info.AppInfo(app)
	commands.Flags(app)

	// users
	user.UserCommand(app)
	user.UserListSubCommand(app)
	user.UserRegisterSubCommand(app)

	// contexts
	context.ContextCommand(app)
	context.ContextListSubCommand(app)

	// applications
	application.ApplicationCommand(app)
	application.ApplicationListSubCommand(app)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
