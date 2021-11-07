package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func PostForm(w http.ResponseWriter, r *http.Request) {
	// parse and get form values
	// err := r.ParseForm()
	// if err != nil {
	// 	panic(err)
	// }
	// fName := r.PostForm.Get("firstName")
	// lName := r.PostForm.Get("lastName")

	// automatically parse form data and get its values
	fName := r.PostFormValue("firstName")
	lName := r.PostFormValue("lastName")
	fmt.Fprintf(w, "%s %s", fName, lName)
}

func TestPostForm(t *testing.T) {
	reqBody := strings.NewReader("firstName=Joe&lastName=Mama")
	req := httptest.NewRequest("POST", "/", reqBody)

	// prequisite for POST FORM
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rec := httptest.NewRecorder()

	PostForm(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
