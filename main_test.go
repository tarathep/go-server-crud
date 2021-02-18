package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tarathep/go-server-crud/apis"
	"github.com/tarathep/go-server-crud/model"
	"github.com/tarathep/go-server-crud/router"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	router := router.Router{apis.HelloHandler{&mockDB{}}, apis.TutorialHandler{}}
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
	router := router.Router{apis.HelloHandler{&mockDB{}}, apis.TutorialHandler{}}
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

// ------ UNIT TEST TUTORIAL API --------

// MOCKS INTERFACE
func (db *mockDB) Create(tutorial model.Tutorial) error {
	return nil
}

func (db *mockDB) FindAll(title string) ([]*model.Tutorial, error) {
	return []*model.Tutorial{
		{
			ID:          primitive.NilObjectID,
			Title:       "TitleTest",
			Description: "DescTest",
			Published:   true,
			CreatedAt:   time.Date(2021, time.February, 2, 1, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2021, time.February, 2, 1, 0, 0, 0, time.UTC),
		},
	}, nil
}

func (db *mockDB) FindOne(id string) (model.Tutorial, error) {
	return model.Tutorial{
		ID:          primitive.NilObjectID,
		Title:       "TitleTest",
		Description: "DescTest",
		Published:   true,
		CreatedAt:   time.Date(2021, time.February, 2, 1, 0, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2021, time.February, 2, 1, 0, 0, 0, time.UTC),
	}, nil
}

func (db *mockDB) Update(tutorial model.Tutorial) error {
	return nil
}

func (db *mockDB) Delete(id string) error {
	return nil
}

func (db *mockDB) DeleteAll() error {
	return nil
}

func (db *mockDB) FindAllPublished() ([]*model.Tutorial, error) {
	return []*model.Tutorial{
		{
			ID:          primitive.NilObjectID,
			Title:       "TitleTest",
			Description: "DescTest",
			Published:   true,
			CreatedAt:   time.Date(2021, time.February, 2, 1, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2021, time.February, 2, 1, 0, 0, 0, time.UTC),
		},
	}, nil
}

// TESTS
func TestReadTutorials(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/tutorials", nil)
	w := httptest.NewRecorder()

	router.Router{
		HelloAPIs:    apis.HelloHandler{DB: &mockDB{}},
		TutorialAPIs: apis.TutorialHandler{DB: &mockDB{}},
	}.Route().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	//test body is not null
	assert.NotNil(t, w.Body)

	//test body response
	assert.Equal(t, strings.Trim(w.Body.String(), "\n"), `[{"id":"000000000000000000000000","title":"TitleTest","description":"DescTest","published":true,"createdAt":"2021-02-02T01:00:00Z","updatedAt":"2021-02-02T01:00:00Z"}]`)
}

func TestReadTutorial(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/tutorials/000000000000000000000000", nil)
	w := httptest.NewRecorder()

	router.Router{
		HelloAPIs:    apis.HelloHandler{DB: &mockDB{}},
		TutorialAPIs: apis.TutorialHandler{DB: &mockDB{}},
	}.Route().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	//test body is not null
	assert.NotNil(t, w.Body)

	//test body response
	assert.Equal(t, strings.Trim(w.Body.String(), "\n"), `{"id":"000000000000000000000000","title":"TitleTest","description":"DescTest","published":true,"createdAt":"2021-02-02T01:00:00Z","updatedAt":"2021-02-02T01:00:00Z"}`)
}
func TestCreateTutorial(t *testing.T) {

	req, _ := http.NewRequest("POST", "/api/tutorials", strings.NewReader(`{
		"title": "xx",
		"description": "xx Description"
	}`))
	w := httptest.NewRecorder()

	router.Router{
		HelloAPIs:    apis.HelloHandler{DB: &mockDB{}},
		TutorialAPIs: apis.TutorialHandler{DB: &mockDB{}},
	}.Route().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	//test body is not null
	assert.NotNil(t, w.Body)

	//test body response
	assert.Equal(t, strings.Trim(w.Body.String(), "\n"), `Inserted a single document Success`)
}

func TestUpdateTutorial(t *testing.T) {
	req, _ := http.NewRequest("PUT", "/api/tutorials/602aa1e04f3b51804eca6917", strings.NewReader(`{"id":"602aa1e04f3b51804eca6917","title":"yy","description":"xx Description","published":false,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}`))
	w := httptest.NewRecorder()

	router.Router{
		HelloAPIs:    apis.HelloHandler{DB: &mockDB{}},
		TutorialAPIs: apis.TutorialHandler{DB: &mockDB{}},
	}.Route().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	//test body is not null
	assert.NotNil(t, w.Body)

	//test body response
	assert.Equal(t, strings.Trim(w.Body.String(), "\n"), `Updated a single document Success`)
}

func TestDeleteTutorial(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/api/tutorials/602aa1e04f3b51804eca6917", nil)
	w := httptest.NewRecorder()

	router.Router{
		HelloAPIs:    apis.HelloHandler{DB: &mockDB{}},
		TutorialAPIs: apis.TutorialHandler{DB: &mockDB{}},
	}.Route().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	//test body is not null
	assert.NotNil(t, w.Body)

	//test body response
	assert.Equal(t, strings.Trim(w.Body.String(), "\n"), `Deleted id:602aa1e04f3b51804eca6917`)
}

func TestDeleteTutorials(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/api/tutorials", nil)
	w := httptest.NewRecorder()

	router.Router{
		HelloAPIs:    apis.HelloHandler{DB: &mockDB{}},
		TutorialAPIs: apis.TutorialHandler{DB: &mockDB{}},
	}.Route().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	//test body is not null
	assert.NotNil(t, w.Body)

	//test body response
	assert.Equal(t, strings.Trim(w.Body.String(), "\n"), `All deleted`)
}
