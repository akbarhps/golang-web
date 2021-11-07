package test

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Golang Auto Escape",
		// auto escape
		// "Body": "<p>Selamat Belajar Golang Web</p>",

		// disable auto escape
		"Body": template.HTML("<p>Selamat Belajar Golang Web</p>"),
	}
	myTemplates.ExecuteTemplate(w, "post.gohtml", data)
}

func TestTemplateAutoEscape(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	TemplateAutoEscape(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	t.Log(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatal(err)
	}
}

func TemplateXSS(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Golang XSS",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	}
	myTemplates.ExecuteTemplate(w, "post.gohtml", data)
}

func TestTemplateXSS(t *testing.T) {
	req := httptest.NewRequest("GET", "/?body=<script>alert(1)</script>", nil)
	rec := httptest.NewRecorder()

	TemplateXSS(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	t.Log(string(body))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatal(err)
	}
}
