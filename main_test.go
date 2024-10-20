package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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

func TestCreateDish(t *testing.T) {
	// Define the request payload
	payload := `{"name": "Test Dish"}`
	req, _ := http.NewRequest("POST", "/dishes", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	// Perform the request using the router set up in TestMain
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	// Check the response status
	assert.Equal(t, http.StatusOK, resp.Code)
	// Check the response body for expected fields
	var response struct {
		Dish struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"dish"`
	}
	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)

	// Check if the response matches the expected values
	assert.Equal(t, "Test Dish", response.Dish.Name)
}
