package main

import (
	"embed"
	"io/fs"

	_ "github.com/joho/godotenv/autoload"
	"github.com/open-function-computers-llc/server-run/server"
)

//go:embed frontend/dist
var frontend embed.FS

func main() {
	// static assets for Angular app
	stripped, err := fs.Sub(frontend, "frontend/dist/frontend")
	if err != nil {
		panic(err)
	}

	s, err := server.New(stripped)
	if err != nil {
		panic(err)
	}
	s.Serve()
}
