package info

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"time"
)

// AppInfo sets the general AppInfo to the CLI
func AppInfo(app *cli.App) {
	// EXAMPLE: Append to an existing template
	cli.AppHelpTemplate = fmt.Sprintf(`%s
WEBSITE: https://axoniq.io/
SUPPORT: support@axoniq.io
`, cli.AppHelpTemplate)

	app.Name = "AxonServer-CLI"
	app.Description = "AxonServer-CLI in GO"
	app.Usage = "This CLI is used to perform actions on AxonServer"
	app.Version = "v1"
	app.Compiled = time.Now()
	app.Copyright = "AxonIQ"
	app.Authors = []*cli.Author{{Name: "Lucas Campos", Email: "lfgcampos@axoniq.io"}}
	app.EnableBashCompletion = true
}
