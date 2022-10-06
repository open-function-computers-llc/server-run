package server

import (
	"errors"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/open-function-computers-llc/server-run/website"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger      *logrus.Logger
	filesystem  fs.FS
	port        int
	sites       []website.Site
	sessions    map[string]session
	adminUser   string
	adminPass   string
	scriptsRoot string
}

func New(filesystem fs.FS) (*Server, error) {
	s := Server{
		logger:     logrus.New(),
		sessions:   map[string]session{},
		filesystem: filesystem,
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return &s, err
	}
	s.port = port

	s.scriptsRoot = os.Getenv("SCRIPTS_ROOT")
	if s.scriptsRoot == "" || s.scriptsRoot[len(s.scriptsRoot)-1:] != "/" {
		return &s, errors.New("Invalid location for SCRIPTS_ROOT. Don't forget the trailing slash.")
	}
	_, err = os.Stat(s.scriptsRoot)
	if os.IsNotExist(err) {
		return &s, errors.New("SCRIPTS_ROOT directory does not exist or is not readable: " + err.Error())
	}

	// TODO: verify all scripts that should exist in the scriptsRoot indeed do exist...

	s.adminUser = os.Getenv("AUTH_USER")
	s.adminPass = os.Getenv("AUTH_PASSWORD")
	if len(s.adminPass) < 8 || s.adminUser == "" {
		return &s, errors.New("Invalid admin user/pass credentials. Password must be at least 8 characters.")
	}

	s.bindRoutes()

	err = s.bootstrapSites()
	if err != nil {
		return &s, err
	}

	return &s, nil
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":"+strconv.Itoa(s.port), nil)
}

func (s *Server) bootstrapSites() error {
	s.sites = []website.Site{}

	stdout, err := exec.Command("ls", os.Getenv("WEBSITES_ROOT")).Output()
	if err != nil {
		return err
	}
	lines := strings.Split(string(stdout), "\n")
	for _, domain := range lines {
		if len(domain) == 0 { // skip blank lines coming from the `ls` command
			continue
		}
		site, err := website.New(domain)
		if err != nil {
			return err
		}
		err = site.LoadStatus()
		if err != nil {
			return err
		}

		s.sites = append(s.sites, site)
	}

	return nil
}

func (s *Server) findAccountByName(a string) (website.Site, error) {
	for _, website := range s.sites {
		if website.Account == a {
			return website, nil
		}
	}

	// no website was found
	return website.Site{}, errors.New("Invalid Account Name")
}
