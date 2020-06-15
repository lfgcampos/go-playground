# Axon Server Command line interface

The intention of this repo is to have a similar cli as in [axon-server-cli](https://github.com/AxonIQ/axon-server-se/tree/master/axonserver-cli) but written in go instead of java.
This is more of a learning exercise of go language but should be as much usable as the official cli.

The Axon Server command line interface allows updating the Axon Server configuration through scripts or from a command line.

For the [Axon Server Standard edition](https://axoniq.io/product-overview/axon-server) the only supported commands are:

* [ ] metrics
* [x] users 
* [x] register-user
* [ ] delete-user

[Axon Server Enterprise edition](https://axoniq.io/product-overview/axon-server-enterprise) supports these additional commands:‌

* [x] applications
* [ ] register-application
* [ ] delete-application
* [ ] init-cluster
* [ ] cluster
* [ ] register-node
* [ ] unregister-node
* [x] contexts
* [ ] register-context
* [ ] delete-context
* [ ] add-node-to-context
* [ ] delete-node-from-context

The option `-S` with the url to the Axon Server is optional, if it is omitted it defaults to [http://localhost:8024](http://localhost:8024/).

## Access control

When running Axon Server with access control enabled, executing commands remotely requires an access token. 
This has to provided with the `-t` option. When you run a command on the Axon Server node itself, you don't have to provide 
a token.

For Axon Server Standard Edition the token is specified in the `axonserver.properties` file \(property name = `axoniq.axonserver.token`\). The token needs to be supplied using the `-t` option in any of the commands.
For this specific cli, you can also set a ENV variable named `AXON_TOKEN` with the desired value or a file named `token` on the same directory of the cli.

## Commands

This section describes all commands supported by the command line interface, grouped by the specific area.
Mind that the list above is marking the ones which are already done.
All the commands have the `-h` option, which will show all the info you need to know including all the flags you can set.

For example:

`./io.axoniq.cli.exe -h`
```
NAME:
   AxonServer-CLI - This CLI is used to perform actions on AxonServer

USAGE:
   io.axoniq.cli.exe [global options] command [command options] [arguments...]

VERSION:
   v1

DESCRIPTION:
   AxonServer-CLI in GO

AUTHOR:
   Lucas Campos <lfgcampos@axoniq.io>

COMMANDS:
   user, u         commands related to users
   context, c      commands related to contexts
   application, a  commands related to applications
   help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --server value, -S value  URL of AxonServer (default: "http://localhost:8024")
   --token value, -t value   Authentication Token [%AXON_TOKEN%]
   --help, -h                show help (default: false)
   --version, -v             print the version (default: false)

COPYRIGHT:
   AxonIQ

WEBSITE: https://axoniq.io/
SUPPORT: support@axoniq.io
```

or
`./io.axoniq.cli.exe user -h`
```
NAME:
   AxonServer-CLI user - commands related to users

USAGE:
   AxonServer-CLI user command [command options] [arguments...]

DESCRIPTION:
   This is the command related to users

COMMANDS:
   list, l      list all the users
   register, r  Register a user
   help, h      Shows a list of commands or help for one command

OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

```

or even deeper
`./io.axoniq.cli.exe user list -h`
```
NAME:
   AxonServer-CLI user list - list all the users

USAGE:
   AxonServer-CLI user list [command options] [arguments...]

DESCRIPTION:
   use to list all the users on the server

OPTIONS:
   --help, -h  show help (default: false)

```