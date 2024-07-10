package commands

import (
	"errors"
	"fmt"

	"github.com/open-function-computers-llc/server-run/website"
	"github.com/urfave/cli"
)

func SetEnv(c *cli.Context) error {
	a := c.String("account")
	site, err := website.FindExistingSite(a)
	if err != nil {
		fmt.Println("error finding existing site")
		return err
	}

	if c.String("database") != "" {
		for _, db := range site.Databases {
			if db.Name == c.String("database") {
				fmt.Println("export DB_ACCOUNT='" + site.Account + "'")
				fmt.Println("export DB_USER='" + db.Username + "'")
				fmt.Println("export DB_NAME='" + db.Name + "'")
				fmt.Println("export DB_PASSWORD='" + db.Password + "'")
				fmt.Println("export DB_HOST='" + db.Host + "'")
				return nil
			}
		}
		return errors.New("Couldn't find database " + c.String("database") + " in account " + site.Account)
	}

	return errors.New("WARNING: no valid flags set!")
}
