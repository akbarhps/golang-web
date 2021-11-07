package test

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	dir := http.Dir("resources")
	fileServer := http.FileServer(dir)

	mux := http.NewServeMux()
	prefix := http.StripPrefix("/static", fileServer)
	mux.Handle("/static/", prefix)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources
var resources embed.FS

func TestFileServerGoEmbed(t *testing.T) {
	dir, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(dir))

	mux := http.NewServeMux()
	prefix := http.StripPrefix("/static", fileServer)
	mux.Handle("/static/", prefix)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
