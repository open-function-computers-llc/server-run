package main

import (
	"embed"
	"io/fs"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/open-function-computers-llc/server-run/commands"
	"github.com/open-function-computers-llc/server-run/server"
	"github.com/urfave/cli"
)

//go:embed frontend/dist
var frontend embed.FS
var Version = "latest-dev"

func main() {
	// check global env file
	_, err := os.Stat("/etc/server-run.env")
	if err == nil {
		// global env file exists! let's load it
		godotenv.Load("/etc/server-run.env")
	}

	app := cli.NewApp()
	app.Name = "OFC Server Run"
	app.Usage = "Manage the website hosting accounts on this server"
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:    "serve",
			Aliases: []string{"s"},
			Usage:   "Start the web based UI",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config, c",
					Usage: "Load configuration from `FILE`",
				},
			},
			Action: func(c *cli.Context) error {
				return startServer()
			},
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "Update a specific account's settings file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "account",
					Usage:    "Account you're updating manually (required)",
					Required: true,
				},
				cli.BoolFlag{
					Name:  "locked",
					Usage: "Locked status for this account",
				},
				cli.StringFlag{
					Name:  "set-domain",
					Usage: "Update the main domain for the selected account",
				},
				cli.StringFlag{
					Name:  "add-domain",
					Usage: "Add a domain to the selected account",
				},
				cli.StringFlag{
					Name:  "remove-domain",
					Usage: "Remove a domain from the selected account",
				},
			},
			Action: func(c *cli.Context) error {
				return commands.Update(c)
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func startServer() error {
	// static assets for Angular app
	stripped, err := fs.Sub(frontend, "frontend/dist/frontend")
	if err != nil {
		return err
	}

	s, err := server.New(stripped)
	if err != nil {
		return err
	}
	return s.Serve()
}
