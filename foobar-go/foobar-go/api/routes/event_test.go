package routes

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestNewEventRoute_CreateEvent(t *testing.T) {

	successString := `{"success":true}`
	eventName := `test event`
	eventLocationId := 2
	eventDatetime := 1667404666

	var jsonData = []byte(`{
		"name": "` + eventName + `",
		"location_id": ` + strconv.Itoa(eventLocationId) + `,
		"datetime": ` + strconv.Itoa(eventDatetime) + `
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/events/", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, successString, w.Body.String())
}

func TestNewEventRoute_GetEventList(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/events/?page_number=0&page_size=10", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
