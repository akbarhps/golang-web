package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/layout.gohtml",
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
	))

	data := map[string]interface{}{
		"Name": r.URL.Query().Get("name"),
	}

	// without template name
	// t.ExecuteTemplate(w, "layout.gohtml", data)

	// with template name
	t.ExecuteTemplate(w, "layout", data)
}

func TestTemplateLayout(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?name=Joe%20Mama", nil)
	rec := httptest.NewRecorder()

	TemplateLayout(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
