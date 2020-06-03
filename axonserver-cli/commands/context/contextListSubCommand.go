package context

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io.axoniq.cli/commands"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// List all contexts
func ContextListSubCommand(app *cli.App) {
	var listContextSubCommand = []*cli.Command{
		{
			Name:        "list",
			Aliases:     []string{"l"},
			Usage:       "list all the contexts",
			Description: "use to list all the contexts on the server",
			Action: func(c *cli.Context) error {
				listContexts()
				return nil
			},
		},
	}
	app.Command("context").Subcommands = append(app.Command("context").Subcommands, listContextSubCommand...)
}

func listContexts() {
	log.Println("calling: " + commands.Server + listContextUrl)
	req, err := http.NewRequest("GET", commands.Server+listContextUrl, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	req.Header.Set(commands.AxonTokenKey, commands.Token)
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}
	fmt.Printf("%s\n", body)
}
