package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Name Can't Be Blank!")
	} else {
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}

func TestResponseCode(t *testing.T) {
	req := httptest.NewRequest("GET", "/?name=Joe", nil)
	rec := httptest.NewRecorder()

	ResponseCode(rec, req)

	res := rec.Result()
	if rec.Code != http.StatusOK {
		panic("StatusCode: " + strconv.Itoa(res.StatusCode) + " Status: " + res.Status)
	} else {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
}

func TestResponseCodeInvalidInput(t *testing.T) {
	req := httptest.NewRequest("GET", "/?name=", nil)
	rec := httptest.NewRecorder()

	ResponseCode(rec, req)

	res := rec.Result()
	if rec.Code != http.StatusBadRequest {
		panic("Response Code Should Be 400 Bad Request")
	}

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
