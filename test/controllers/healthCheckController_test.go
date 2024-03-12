package main

import (
	"encoding/json"
	"golang-marketplace-app/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnHttpStatusOk_whenServerIsNormal(t *testing.T) {
	router := router.StartApp()

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/health-check", nil)
	router.ServeHTTP(recorder, request)

	assert.Equal(t, 200, recorder.Code)
	var responseBody map[string]string
	if err := json.NewDecoder(recorder.Body).Decode(&responseBody); err != nil {
			t.Fatal(err)
	}
	expectedMessage := "success"
	actualMessage, ok := responseBody["message"]
	assert.True(t, ok, "Response body does not contain 'message' field")
	assert.Equal(t, expectedMessage, actualMessage)
}