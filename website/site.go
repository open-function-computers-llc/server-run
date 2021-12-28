package website

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Site struct {
	IsLocked         bool     `json:"isLocked"`
	Domain           string   `json:"domain"`
	AlternateDomains []string `json:"alternateDomains"`
	UptimeURI        string   `json:"uptimeURI"`
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

	verifyExists(s.Domain, "settings.json")
	err := s.hydrateData()

	return err
}

func verifyExists(folderName, fileName string) {
	// check folder
	folderPath := os.Getenv("CONFIG_FILE_ROOT") + "/" + folderName
	_, err := os.Stat(folderPath)
	if err != nil {
		os.Mkdir(folderPath, 0755)
	}

	// check file
	_, err = os.Stat(folderPath + "/" + fileName)
	if err != nil {
		data := settings{
			Domain: folderName,
			AlwaysUnlockedDirectories: []string{
				os.Getenv("WEBSITES_ROOT") + "/" + folderName + "/uploads", // uploads directory is always flagged as g2g
			},
		}

		// TODO: juggle the different server types to hydrate the default unlocked directories

		json, _ := json.Marshal(data)
		file, _ := os.Create(folderPath + "/" + fileName)
		file.WriteString(string(json))
	}
}

func (s *Site) hydrateData() error {
	directory := os.Getenv("CONFIG_FILE_ROOT")
	status, err := ioutil.ReadFile(directory + "/" + s.Domain + "/settings.json")
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
