package test

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request")
		return
	}

	// change Content-Disposition to force download
	// default Content-Disposition is inline
	// inline means that the file will be displayed in browser
	w.Header().Add("Content-Disposition", "attachment; filename="+fileName+"\"")
	http.ServeFile(w, r, "./resources/"+fileName)
}

func TestDownloadFileServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", DownloadFile)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Error(err)
	}
}
