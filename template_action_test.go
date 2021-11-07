package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"text/template"
)

func TemplateIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	data := map[string]interface{}{
		"Name": r.URL.Query().Get("name"),
	}

	t.ExecuteTemplate(w, "if.gohtml", data)
}

func TestTemplateIf(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?name=Joe%20Mama", nil)
	rec := httptest.NewRecorder()

	TemplateIf(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TemplateOperator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))

	// need to cast to int so the template can compare
	score, _ := strconv.Atoi(r.URL.Query().Get("score"))
	data := map[string]interface{}{
		"Score": score,
	}

	t.ExecuteTemplate(w, "comparator.gohtml", data)
}

func TestTemplateOperator(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?score=80", nil)
	rec := httptest.NewRecorder()

	TemplateOperator(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TemplateRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	data := map[string]interface{}{
		"Items": []string{"Joe", "Mama", "Thicc"},
	}

	t.ExecuteTemplate(w, "range.gohtml", data)
}

func TestTemplateRange(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	TemplateRange(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TemplateWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))

	data := map[string]interface{}{
		"Name": r.URL.Query().Get("name"),
		"Address": map[string]string{
			"Street": "Jalan Surga Gang Neraka",
			"City":   "Akhirat",
		},
	}

	t.ExecuteTemplate(w, "with.gohtml", data)
}

func TestTemplateWith(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?name=Joe%20Mama", nil)
	rec := httptest.NewRecorder()

	TemplateWith(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
