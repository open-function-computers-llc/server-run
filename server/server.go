package server

import (
	"io/fs"
	"net/http"
	"os"
	"strconv"

	"github.com/go-kit/kit/log/logrus"
)

type Server struct {
	logger     *logrus.Logger
	filesystem fs.FS
	port       int
}

func New(filesystem fs.FS) (*Server, error) {
	s := Server{}
	s.filesystem = filesystem

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return &s, err
	}
	s.port = port

	s.bindRoutes()

	return &s, nil
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":"+strconv.Itoa(s.port), nil)
}
