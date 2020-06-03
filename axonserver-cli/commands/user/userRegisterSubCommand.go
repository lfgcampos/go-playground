package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"io.axoniq.cli/commands"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Register a user
func UserRegisterSubCommand(app *cli.App) {
	var registerUserSubCommand = []*cli.Command{
		{
			Name:        "register",
			Aliases:     []string{"r"},
			Usage:       "Register a user",
			Description: "register a user to be used on axonserver",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "username",
					Usage:       "user username",
					Aliases:     []string{"u"},
					Required:    true,
					Destination: &username,
				},
				&cli.StringFlag{
					Name:        "password",
					Usage:       "user password",
					Aliases:     []string{"p"},
					Required:    true,
					Destination: &password,
				},
				&cli.StringSliceFlag{
					Name:        "roles",
					Usage:       "user roles",
					Aliases:     []string{"r"},
					Destination: &roles,
				},
			},
			Action: func(c *cli.Context) error {
				registerUser()
				return nil
			},
		},
	}
	app.Command("user").Subcommands = append(app.Command("user").Subcommands, registerUserSubCommand...)
}

func registerUser() {
	log.Println("calling: " + commands.Server + registerUserUrl)
	userJson := buildUserJson()
	req, err := http.NewRequest("POST", commands.Server+registerUserUrl, bytes.NewBuffer(userJson))
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	req.Header.Set(commands.AxonTokenKey, commands.Token)
	req.Header.Set(commands.ContentType, commands.JsonType)
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

func buildUserJson() []byte {
	user := &user{
		Username: username,
		Password: password,
		Roles:    roles.Value(),
	}
	userJson, err := json.Marshal(&user)
	if err != nil {
		log.Fatal("Error building the user json. ", err)
	}
	fmt.Printf("userJson: %+v\n", string(userJson))
	return userJson
}
