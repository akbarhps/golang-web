package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func TestHTTP(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "https://localhost:8080/hello", nil)
	rec := httptest.NewRecorder()

	HelloHandler(rec, req)

	result := rec.Result()
	body, _ := io.ReadAll(result.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
