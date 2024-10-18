package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	router = SetupApp()
	m.Run()
}

func TestPingRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Pong Uwu", w.Body.String())
}

func TestGetRestaurants(t *testing.T) {
	//Test not passing in "search_term" to get 400 erro
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/restaurants", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

	// Test passing "search_term" = "Milano"
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/restaurants?search_term=Milano", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response struct {
		Restaurants []struct {
			Name string `json:"Name"`
		} `json:"restaurants"`
	}

	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)

	//t.Log(response.Restaurants)

	// Assert that each restaurant's name contains "Milano"
	for _, restaurant := range response.Restaurants {
		assert.Contains(t, strings.ToLower(restaurant.Name), "milano")
	}
}
