package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "X-USERNAME",
		Value: r.URL.Query().Get("name"),
	}

	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "Cookie is set ", cookie)
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-USERNAME")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Fprintf(w, "Cookie not found!")
	} else {
		fmt.Fprintf(w, "Welcome Back %s", cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", func(w http.ResponseWriter, r *http.Request) {
		SetCookie(w, r)
	})
	mux.HandleFunc("/get-cookie", func(w http.ResponseWriter, r *http.Request) {
		GetCookie(w, r)
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "/?name=John", nil)
	rec := httptest.NewRecorder()

	SetCookie(rec, req)

	cookie := rec.Result().Cookies()
	fmt.Println(cookie)
}

func TestGetCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	cookie := http.Cookie{
		Name:  "X-USERNAME",
		Value: "John",
	}

	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	GetCookie(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TestGetCookieInvalid(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	GetCookie(rec, req)

	res := rec.Result()
	if res.StatusCode != http.StatusBadRequest {
		panic("Status Code Should Be 404 Bad Request!")
	}

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
