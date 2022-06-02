package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestJsonResponce(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api", nil)
	router.ServeHTTP(w, req)

	var response map[string]int32
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	_, existsA := response["a"]
	_, existsB := response["b"]

	assert.Nil(t, err)
	assert.True(t, existsA)
	assert.True(t, existsB)
}
