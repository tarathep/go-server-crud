package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tarathep/go-server-crud/apis"
	"github.com/tarathep/go-server-crud/model"
	"github.com/tarathep/go-server-crud/router"
)

type mockDB struct{}

func (db *mockDB) AllHello() ([]*model.Hello, error) {
	hellos := make([]*model.Hello, 0)
	hellos = append(hellos, &model.Hello{"C++", "c is height lv programing"})
	hellos = append(hellos, &model.Hello{"VB", "c is basic programing"})

	return hellos, nil
}

func (db *mockDB) InsertHello(hello model.Hello) (model.Hello, error) {
	return hello, nil
}

func TestGetHello(t *testing.T) {
	router := router.Router{apis.HelloHandler{&mockDB{}}}
	r := router.Route()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	//test body is not null
	assert.NotNil(t, w.Body)

	//test body response
	assert.Equal(t, strings.Trim(w.Body.String(), "\n"), `[{"language":"C++","message":"c is height lv programing"},{"language":"VB","message":"c is basic programing"}]`)
}

func TestHelloAdd(t *testing.T) {
	router := router.Router{apis.HelloHandler{&mockDB{}}}
	r := router.Route()

	w := httptest.NewRecorder()
	inputJSON := `{"language":"C++","message":"c is height lv programing"}`
	req, _ := http.NewRequest("POST", "/hello", strings.NewReader(inputJSON))
	r.ServeHTTP(w, req)
	//test http status ok = 200
	assert.Equal(t, http.StatusOK, w.Code)

	//test body is not null
	assert.NotNil(t, w.Body)

	//test body response
	assert.Equal(t, strings.Trim(w.Body.String(), "\n"), "success")

}
