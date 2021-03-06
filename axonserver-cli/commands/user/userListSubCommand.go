package user

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
func UserListSubCommand(app *cli.App) {
	var listUserSubCommand = []*cli.Command{
		{
			Name:        "list",
			Aliases:     []string{"l"},
			Usage:       "list all the users",
			Description: "use to list all the users on the server",
			Action: func(c *cli.Context) error {
				listUsers()
				return nil
			},
		},
	}
	app.Command("user").Subcommands = append(app.Command("user").Subcommands, listUserSubCommand...)
}

func listUsers() {
	log.Println("calling: " + commands.Server + listUserUrl)
	req, err := http.NewRequest("GET", commands.Server+listUserUrl, nil)
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
