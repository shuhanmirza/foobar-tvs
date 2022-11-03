package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http/httptest"
	"testing"
)

// TODO: write proper test
func TestEventService_GetEventById(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	event, err := eventService.GetEventById(ctx, 1)

	log.Println(err)
	log.Println(event)
}
