package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type Page struct {
	Title string
	Name  string
}

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/name.gohtml"))

	data := map[string]interface{}{
		"Title": "Template Dari Map",
		"Name":  "Joe Mama",
	}

	t.ExecuteTemplate(w, "name.gohtml", data)
}

func TestTemplateDataMap(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	TemplateDataMap(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/name.gohtml"))

	data := Page{
		Title: "Template Dari Map",
		Name:  "Joe Mama",
	}

	t.ExecuteTemplate(w, "name.gohtml", data)
}

func TestTemplateDataStruct(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	TemplateDataStruct(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
