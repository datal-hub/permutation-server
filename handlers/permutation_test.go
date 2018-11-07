package handlers_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	. "permutation-server/handlers"
	"permutation-server/models"
	testData "permutation-server/models/testing"
	"permutation-server/pkg/database"
)

func TestInitNilBody(t *testing.T) {
	database.Testing = true
	var req = httptest.NewRequest("POST", "http://perm-server.ru/v1/init", nil)
	resp := httptest.NewRecorder()
	Init(resp, req)
	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, resp.Code)
}

func TestInit(t *testing.T) {
	database.Testing = true
	body, _ := json.Marshal([]int64{1, 2, 3})
	var req = httptest.NewRequest("POST", "http://perm-server.ru/v1/init", bytes.NewReader(body))
	resp := httptest.NewRecorder()
	Init(resp, req)
	assert := assert.New(t)
	assert.Equal(http.StatusCreated, resp.Code)
}

func TestNext(t *testing.T) {
	database.Testing = true
	var req = httptest.NewRequest("GET", "http://perm-server.ru/v1/next", nil)
	req.AddCookie(
		&http.Cookie{
			Name:  "uid",
			Value: testData.TestPermutation.Uuid,
			Path:  "/"})
	resp := httptest.NewRecorder()
	Next(resp, req)
	b, _ := ioutil.ReadAll(resp.Body)
	arr := make([]int64, 0)
	err := json.Unmarshal(b, &arr)
	if err != nil {
		t.Fatal(err)
	}
	res := models.Permutation{
		Uuid: testData.TestPermutation.Uuid,
		Data: arr,
	}
	assert := assert.New(t)
	assert.Equal(http.StatusOK, resp.Code)
	assert.Equal(testData.TestPermutation, res)
}
