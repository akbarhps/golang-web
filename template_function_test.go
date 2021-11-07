package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name is " + myPage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Joe"}}`))

	data := MyPage{
		Name: "Joe Mama",
	}

	t.ExecuteTemplate(w, "FUNCTION", data)
}

func TestTemplateFunction(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	TemplateFunction(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))

	t.ExecuteTemplate(w, "FUNCTION", map[string]interface{}{
		"Name": "Joe Mama",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	TemplateFunctionGlobal(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TemplateCreateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{upper .Name}}`))
	t.ExecuteTemplate(w, "FUNCTION", map[string]interface{}{
		"Name": "Joe Mama",
	})
}

func TestTemplateCreateFunctionGlobal(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	TemplateCreateFunctionGlobal(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TemplateFunctionPipeline(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(value string) string {
			return "Hello " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{sayHello .Name | upper}}`))
	t.ExecuteTemplate(w, "FUNCTION", map[string]interface{}{
		"Name": "Joe Mama",
	})
}

func TestTemplateFunctionPipeline(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	TemplateFunctionPipeline(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
