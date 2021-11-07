package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before handler execution")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After handler execution")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered From Error: ", err)

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal Server Error: %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(w, r)
}

func TestMiddlewareLog(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Log("Handler Executed")
		fmt.Fprint(w, "Hello, World!")
	})

	middleware := &LogMiddleware{
		Handler: mux,
	}

	server := httptest.NewServer(middleware)
	defer server.Close()

	req := httptest.NewRequest("GET", server.URL+"/", nil)
	rec := httptest.NewRecorder()

	middleware.ServeHTTP(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	defer rec.Result().Body.Close()

	t.Log(string(body))
}

func TestMiddlewareErrorHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Something went wrong")
	})

	middleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: middleware,
	}

	server := httptest.NewServer(errorHandler)
	defer server.Close()

	t.Log(server.URL)

	req := httptest.NewRequest("GET", server.URL+"/panic", nil)
	rec := httptest.NewRecorder()

	errorHandler.ServeHTTP(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	defer rec.Result().Body.Close()

	t.Log(string(body))
}
