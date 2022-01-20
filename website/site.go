package website

import (
	"encoding/json"
	"errors"
	"os"
)

var stateFilename = "settings.json"

type Site struct {
	IsLocked         bool     `json:"isLocked"`
	Domain           string   `json:"domain"`
	AlternateDomains []string `json:"alternateDomains"`
	UptimeURI        string   `json:"uptimeURI"`
	Username         string   `json:"username"`
}

func New(domain string) (Site, error) {
	s := Site{
		Domain:           domain,
		AlternateDomains: []string{},
	}

	return s, nil
}

// LoadStatus - hydrate a site by the contents of of a JSON config file
func (s *Site) LoadStatus() error {
	if s.Domain == "" {
		return errors.New("Can't load status for an empty domain")
	}

	s.verifyStateFileExists()
	return s.hydrateData()
}

func (s *Site) stateFolder() string {
	return os.Getenv("ACCOUNTS_ROOT") + s.Domain
}

func (s *Site) stateFilePath() string {
	return s.stateFolder() + "/" + stateFilename
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
		data := settings{
			Domain: s.Domain,
			AlwaysUnlockedDirectories: []string{
				os.Getenv("WEBSITES_ROOT") + s.Domain + "/uploads", // uploads directory is always flagged as g2g
			},
		}

		// TODO: juggle the different server types to hydrate the default unlocked directories

		json, _ := json.Marshal(data)
		file, _ := os.Create(s.stateFilePath())
		file.WriteString(string(json))
	}
}

func (s *Site) saveStateFile() error {
	json, _ := json.Marshal(s)
	return os.WriteFile(s.stateFilePath(), json, 0644)
}

func (s *Site) hydrateData() error {
	directory := os.Getenv("ACCOUNTS_ROOT")
	status, err := os.ReadFile(directory + s.Domain + "/settings.json")
	if err != nil {
		return err
	}
	var siteSettings settings
	err = json.Unmarshal(status, &siteSettings)
	if err != nil {
		return err
	}

	// now we hydrate the data from the json
	s.IsLocked = siteSettings.IsLocked
	s.UptimeURI = siteSettings.UptimeURI

	return nil
}
