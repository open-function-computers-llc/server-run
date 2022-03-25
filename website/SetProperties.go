package website

import (
	"errors"
	"strings"
)

func (ws *Site) SetUptimeURI(uri string) error {
	ws.UptimeURI = uri
	return ws.saveStateFile()
}

func (ws *Site) SetLocked(locked bool) error {
	ws.IsLocked = locked
	return ws.saveStateFile()
}

func (ws *Site) SetPrimaryDomain(d string) error {
	ws.PrimaryDomain = d
	return ws.saveStateFile()
}

func (ws *Site) AddAlternateDomain(d string) error {
	d = strings.TrimSpace(d) // trim it down

	// check to see if this matches the primary domain
	if d == ws.PrimaryDomain {
		return errors.New("You can't add the primary domain to the list of alternate domains")
	}

	// check to see if the domain is already in the list
	for _, existingDomain := range ws.AlternateDomains {
		if existingDomain == d {
			return errors.New("Can not add domain " + d + " to alternate domain list, it's already in the list")
		}
	}

	ws.AlternateDomains = append(ws.AlternateDomains, d)
	return ws.saveStateFile()
}

func (ws *Site) RemoveAlternateDomain(d string) error {
	d = strings.TrimSpace(d) // trim it down

	// build a new list and skip the current one if it's in there already
	newDomainList := []string{}
	for _, existingDomain := range ws.AlternateDomains {
		if existingDomain == d {
			continue
		}
		newDomainList = append(newDomainList, existingDomain)
	}

	ws.AlternateDomains = newDomainList
	return ws.saveStateFile()
}
