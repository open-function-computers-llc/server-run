package website

// this is the structure of the JSON file that we will use for configuration.
// make sure this matches the bash scripts for reading and writting settings.
type settings struct {
	IsLocked                  bool     `json:"isLocked"`
	UptimeURI                 string   `json:"uptimeURI"`
	Domain                    string   `json:"domain"`
	AlwaysUnlockedDirectories []string `json:"alwaysUnlockedDirectories"`
	AlternateDomains          []string `json:"alternateDomains"`
}
