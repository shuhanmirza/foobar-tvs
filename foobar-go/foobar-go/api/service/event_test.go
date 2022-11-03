package service

import (
	"foobar-go/api/structs"
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

// TODO: write proper test
func TestEventService_GetEventListByPage(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	response, err := eventService.GetEventListByPage(ctx, structs.GetEventListByPageRequest{
		PageNumber: 0,
		PageSize:   0,
	})

	log.Println(response)
	log.Println(err)

}
