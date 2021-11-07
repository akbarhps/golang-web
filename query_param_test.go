package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func TestQueryParam(t *testing.T) {
	req := httptest.NewRequest("GET", "/?name=Stipen", nil)
	rec := httptest.NewRecorder()

	SayHello(rec, req)

	resp := rec.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func MultipleQueryParameter(w http.ResponseWriter, r *http.Request) {
	fName := r.URL.Query().Get("first_name")
	lName := r.URL.Query().Get("last_name")
	fmt.Fprintf(w, "Hello, %s %s!", fName, lName)
}

func TestMultipleQueryParameter(t *testing.T) {
	req := httptest.NewRequest("GET", "/?first_name=Stipen&last_name=Debandits", nil)
	rec := httptest.NewRecorder()

	MultipleQueryParameter(rec, req)

	resp := rec.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func MultipleValueQueryParameter(w http.ResponseWriter, r *http.Request) {
	favorite := r.URL.Query()["favorite"]
	fmt.Fprintf(w, "Your favorite food is %s\n", strings.Join(favorite, " "))
}

func TestMultipleValueQueryParameter(t *testing.T) {
	req := httptest.NewRequest("GET", "/?favorite=Pizza&favorite=Burger", nil)
	rec := httptest.NewRecorder()

	MultipleValueQueryParameter(rec, req)

	resp := rec.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
