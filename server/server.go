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
	logger     *logrus.Logger
	filesystem fs.FS
	port       int
	sites      []website.Site
	sessions   map[string]session
	adminUser  string
	adminPass  string
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

	s.adminUser = os.Getenv("AUTHUSER")
	s.adminPass = os.Getenv("AUTHPASSWORD")

	s.bindRoutes()
	err = s.bootstrapSites()
	if err != nil {
		return &s, err
	}

	if s.adminPass == "" || s.adminUser == "" {
		return &s, errors.New("AUTHUSER and AUTHPASSWORD are required env vars")
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
