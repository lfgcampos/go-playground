package application

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io.axoniq.cli/commands"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// List all users
func ApplicationListSubCommand(app *cli.App) {
	var listApplicationSubCommand = []*cli.Command{
		{
			Name:        "list",
			Aliases:     []string{"l"},
			Usage:       "list all the applications",
			Description: "use to list all the applications on the server",
			Action: func(c *cli.Context) error {
				listApplications()
				return nil
			},
		},
	}
	app.Command("application").Subcommands = append(app.Command("application").Subcommands, listApplicationSubCommand...)
}

func listApplications() {
	log.Println("calling: " + commands.Server + listApplicationUrl)
	req, err := http.NewRequest("GET", commands.Server+listApplicationUrl, nil)
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
