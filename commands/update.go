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
	if c.String("locked") != "" {
		if c.String("locked") != "true" && c.String("locked") != "false" {
			return errors.New("Invalid option for 'locked' status. Valid options: true|false.")
		}
		if c.String("locked") == "true" {
			return site.SetLocked(true)
		}
		if c.String("locked") == "false" {
			return site.SetLocked(false)
		}
		return errors.New("something weird just happened")
	}

	if c.String("set-domain") != "" {
		return site.SetPrimaryDomain(c.String("set-domain"))
	}

	if c.String("add-domain") != "" {
		return site.AddAlternateDomain(c.String("add-domain"))
	}

	if c.String("remove-domain") != "" {
		return site.RemoveAlternateDomain(c.String("remove-domain"))
	}
	return errors.New("WARNING: No update command was ran!")
}
