package website

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

var stateFilename = "settings.json"

type Site struct {
	IsLocked                  bool     `json:"isLocked"`
	Account                   string   `json:"account"`
	UptimeURI                 string   `json:"uptimeURI"`
	Username                  string   `json:"username"`
	PrimaryDomain             string   `json:"domain"`
	AlternateDomains          []string `json:"alternateDomains"`
	AlwaysUnlockedDirectories []string `json:"alwaysUnlockedDirectories"`
	PubKey                    string   `json:"sshPubKey"`
}

func New(account string) (Site, error) {
	s := Site{
		Account:          account,
		AlternateDomains: []string{},
	}

	return s, nil
}

// LoadStatus - hydrate a site by the contents of of a JSON config file and
// create the config file if it is missing
func (s *Site) LoadStatus() error {
	if s.Account == "" {
		return errors.New("Can't load status for an empty domain")
	}

	s.verifyStateFileExists()
	return s.hydrateData()
}

// LoadStatus - hydrate a site by the contents of of a JSON config file
func (s *Site) loadExistingStatus() error {
	if s.Account == "" {
		return errors.New("Can't load status for an empty domain")
	}

	return s.loadDataFromExistingStateFile()
}

func (s *Site) stateFolder() string {
	return os.Getenv("ACCOUNTS_ROOT") + s.Account
}

func (s *Site) stateFilePath() string {
	return s.stateFolder() + "/" + stateFilename
}

func (s *Site) hydrateSSHPubKey() {
	// check for key file
	_, err := os.Stat(s.stateFolder() + "/.ssh/id_rsa.pub")
	if err != nil {
		log.Println(err)
		return
	}

	bytes, err := os.ReadFile(s.stateFolder() + "/.ssh/id_rsa.pub")
	if err != nil {
		log.Println(err)
		return
	}

	s.PubKey = string(bytes)
}

func (s *Site) verifyStateFileExists() {
	// check folder
	_, err := os.Stat(s.stateFolder())
	if err != nil {
		os.Mkdir(s.stateFolder(), 0755)
	}

	// check file
	_, err = os.Stat(s.stateFilePath())
	if err != nil {
		s.AlwaysUnlockedDirectories = []string{
			os.Getenv("WEBSITES_ROOT") + s.Account + "/uploads", // uploads directory is always flagged as g2g
		}
		s.Username = s.Account

		// TODO: juggle the different server types to hydrate the default unlocked directories

		json, _ := json.Marshal(s)
		file, _ := os.Create(s.stateFilePath())
		file.WriteString(string(json))
	}
}

func (s *Site) loadDataFromExistingStateFile() error {
	// check folder
	_, err := os.Stat(s.stateFolder())
	if err != nil {
		fmt.Println("bad folder " + s.stateFolder())
		return err
	}

	// check file
	_, err = os.Stat(s.stateFilePath())
	if err != nil {
		fmt.Println("settings file doesn't exist: " + s.stateFilePath())
		fmt.Println("creating it...")
		s.verifyStateFileExists() // create the file
	}
	return s.hydrateData()
}

func (s *Site) saveStateFile() error {
	json, _ := json.Marshal(s)
	return os.WriteFile(s.stateFilePath(), json, 0644)
}

func (s *Site) hydrateData() error {
	directory := os.Getenv("ACCOUNTS_ROOT")
	status, err := os.ReadFile(directory + s.Account + "/settings.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(status, &s)
	if err != nil {
		return err
	}

	s.hydrateSSHPubKey()

	return nil
}

func FindExistingSite(d string) (*Site, error) {
	s, err := New(d)
	if err != nil {
		return &s, err
	}
	err = s.loadDataFromExistingStateFile()
	if err != nil {
		fmt.Println("couldn't load existing state file")
		return &s, err
	}
	return &s, err
}
