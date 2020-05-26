package commands

import "github.com/urfave/cli/v2"

var contentType = "Content-Type"
var jsonType = "application/json"

var axonTokenKey = "AxonIQ-Access-Token"
var server = "http://localhost:8024"
var token string
var username, password string
var roles cli.StringSlice

var listContextUrl = "/v1/public/context"
var listApplicationUrl = "/v1/public/applications"

var listUserUrl = "/v1/public/users"
var registerUserUrl = "/v1/users"

type user struct {
	Username string `json:"userName"`
	Password string `json:"password"`
	Roles []string `json:"roles"`
}
