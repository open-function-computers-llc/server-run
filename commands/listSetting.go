package commands

import (
	"errors"
	"fmt"

	"github.com/open-function-computers-llc/server-run/website"
	"github.com/urfave/cli"
)

func ListSetting(c *cli.Context) error {
	a := c.String("account")
	site, err := website.FindExistingSite(a)
	if err != nil {
		fmt.Println("error finding existing site")
		return err
	}

	setting := c.String("setting")

	if setting == "databases" {
		for _, db := range site.Databases {
			fmt.Println(db.Name)
		}
		return nil
	}

	return errors.New("WARNING: invalid setting requested")
}
