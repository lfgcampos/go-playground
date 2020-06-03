package user

import "github.com/urfave/cli/v2"

var listUserUrl = "/v1/public/users"
var registerUserUrl = "/v1/users"

var username, password string
var roles cli.StringSlice

type user struct {
	Username string   `json:"userName"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}
