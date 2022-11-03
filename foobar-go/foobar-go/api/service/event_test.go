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

	response, err := eventService.GetEventListByPage(ctx, structs.GetEventListByPageQueryRequest{
		PageNumber: 0,
		PageSize:   0,
	})

	log.Println(response)
	log.Println(err)

}

// TODO: write proper test
func TestEventService_UpdateEvent(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	response, err := eventService.UpdateEvent(ctx, 40, structs.UpdateEventRequest{
		Name:       "test",
		LocationId: 100,
		Datetime:   1667457058,
	})

	log.Println(response)
	log.Println(err)
}

// TODO: write proper tests
func TestEventService_DeleteEvent(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	response, err := eventService.DeleteEvent(ctx, 42)

	log.Println(response)
	log.Println(err)
}
