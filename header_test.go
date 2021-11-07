package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	fmt.Fprintf(w, "Content-Type is %s", contentType)
}

func TestRequestHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	RequestHeader(rec, req)

	resp := rec.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Response Header has been changed!")
}

func TestResponseHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	ResponseHeader(rec, req)

	resp := rec.Result()
	fmt.Println(resp.Header.Get("Content-Type"))

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
