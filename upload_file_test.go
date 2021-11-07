package test

import (
	"bytes"
	_ "embed"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	err := myTemplates.ExecuteTemplate(w, "upload.form.gohtml", nil)
	if err != nil {
		panic(err)
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form,
	// 32 << 20 means that the maximum memory file size that can be stored in memory is 32 megabytes
	// r.ParseMultipartForm(32 << 20)

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	// set uploaded file destination
	fileDest, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	// copy uploaded file to destination
	_, err = io.Copy(fileDest, file)
	if err != nil {
		panic(err)
	}

	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadFormServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/img.png
var image []byte

func TestUploadForm(t *testing.T) {
	// Create a new multipart form (must be the same as the one used in the html)
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// fill form name field
	writer.WriteField("name", "Joe mama")

	// fill form upload file field
	// "file" -> name of the form field
	// "image.png" -> file name to send to the server (this is the name that will be saved on the server)
	file, _ := writer.CreateFormFile("file", "image.png")

	// write embedded file to form upload file field
	file.Write(image)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	rec := httptest.NewRecorder()

	// set request header to multipart form
	req.Header.Set("Content-Type", writer.FormDataContentType())

	Upload(rec, req)

	responseBody, _ := ioutil.ReadAll(rec.Result().Body)
	t.Log(string(responseBody))
}
