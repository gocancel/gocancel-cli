# gocancel-cli

The GoCancel CLI helps you interact with your GoCancel account right from the terminal.

## Usage

Installing the CLI provides access to the `gocancel` command.

```sh
gocancel [command]

# Run `--help` for detailed information about CLI commands
gocancel [command] help
```

## Authenticating with GoCancel

To use `gocancel`, you need to authenticate with GoCancel by providing the client credentials of an application, which can be created from the _Applications_ section of your account. Please use the following scopes when creating:

```
read:categories read:letters write:letters read:organizations
```

You have to set the `GOCANCEL_CLIENT_ID` and `GOCANCEL_CLIENT_SECRET` environmental variables to pass the credentials to the CLI. To override the endpoint used by the underlaying API client, you can set the `GOCANCEL_API_URL`.
