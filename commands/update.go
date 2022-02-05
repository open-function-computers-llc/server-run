package commands

import (
	"errors"
	"fmt"

	"github.com/open-function-computers-llc/server-run/website"
	"github.com/urfave/cli"
)

func Update(c *cli.Context) error {
	a := c.String("account")
	site, err := website.FindExistingSite(a)
	if err != nil {
		fmt.Println("error finding existing site")
		return err
	}

	// now we've loaded a site, check the other flags
	if c.IsSet("locked") {
		return site.SetLocked(c.Bool("locked"))
	}
	return errors.New("WARNING: No update command was ran!")
}
