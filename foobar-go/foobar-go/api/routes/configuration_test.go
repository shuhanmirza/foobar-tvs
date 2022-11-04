package routes

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewConfigurationRoute_Status(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestNewConfigurationRoute_GetLocation(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/config/location-list", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
