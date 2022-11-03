package service

import (
	"database/sql"
	"foobar-go/api/structs"
	"foobar-go/infrastructure"
	db "foobar-go/sqlc"
	"foobar-go/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"log"
)

type EventService struct {
	store *infrastructure.Store
}

func NewEventService(store *infrastructure.Store) EventService {
	return EventService{
		store: store,
	}
}

func (s EventService) CreateEvent(ctx *gin.Context, request structs.CreateEventRequest) (response structs.CreateEventResponse, err error) {
	_, err = s.store.Queries.CreateEvent(ctx, db.CreateEventParams{
		Name:       request.Name,
		LocationID: request.LocationId,
		Datetime:   request.Datetime,
	})

	if err != nil {
		response = structs.CreateEventResponse{Success: false}

		if pgErr, isPGErr := err.(*pq.Error); isPGErr {
			if pgErr.Code == util.SQL_FOREIGN_KEY_VIOLATION_ERROR_CODE {
				return response, &util.InvalidLocationId{}
			}
		}

		log.Println("error while creating event")
		log.Println(err)
		return response, err

	}

	response = structs.CreateEventResponse{Success: true}

	return response, err
}

func (s EventService) GetEventById(ctx *gin.Context, eventId int64) (response structs.GetEventByIdResponse, err error) {
	event, err := s.store.Queries.GetEventWithLocationById(ctx, eventId)
	if err != nil {
		if err == sql.ErrNoRows {
			return response, &util.InvalidEventId{}
		}

		log.Println("error while reading event")
		log.Println(err)
		return response, err
	}

	response = structs.GetEventByIdResponse{
		Name:     event.Name,
		Location: event.Location,
		Datetime: event.Datetime,
	}

	return response, err
}

func (s EventService) GetEventListByPage(ctx *gin.Context, request structs.GetEventListByPageQueryRequest) (response structs.GetEventListByPageResponse, err error) {
	eventList, err := s.store.Queries.GetEventListByOffsetLimit(ctx, db.GetEventListByOffsetLimitParams{
		Offset: request.PageNumber * request.PageSize,
		Limit:  request.PageSize,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			response = structs.GetEventListByPageResponse{Events: nil}
			return response, nil
		}

		log.Println("error while reading events")
		log.Println(err)
		return response, err
	}

	var eventJsonList []structs.EventJson

	for _, event := range eventList {
		eventJsonList = append(eventJsonList, structs.EventJson{
			Id:       event.ID,
			Name:     event.Name,
			Location: event.Location,
			Datetime: event.Datetime,
		})
	}

	response = structs.GetEventListByPageResponse{Events: eventJsonList}
	return response, err
}

func (s EventService) UpdateEvent(ctx *gin.Context, eventId int64, request structs.UpdateEventRequest) (response structs.UpdateEventResponse, err error) {
	_, err = s.store.Queries.UpdateEvent(ctx, db.UpdateEventParams{
		ID:         eventId,
		Name:       request.Name,
		LocationID: request.LocationId,
		Datetime:   request.Datetime,
	})

	if err != nil {
		response = structs.UpdateEventResponse{Success: false}

		if err == sql.ErrNoRows {
			return response, &util.InvalidEventId{}
		}

		if pgErr, isPGErr := err.(*pq.Error); isPGErr {
			if pgErr.Code == util.SQL_FOREIGN_KEY_VIOLATION_ERROR_CODE {
				return response, &util.InvalidLocationId{}
			}
		}

		log.Println("error while updating event")
		log.Println(err)
		return response, err

	}

	response = structs.UpdateEventResponse{Success: true}
	return response, err
}

func (s EventService) DeleteEvent(ctx *gin.Context, eventId int64) (response structs.DeleteEventResponse, err error) {
	_, err = s.store.Queries.DeleteEvent(ctx, eventId)

	if err != nil {
		response = structs.DeleteEventResponse{Success: false}
		if err == sql.ErrNoRows {
			return response, &util.InvalidEventId{}
		}
		log.Println("error while deleting event")
		log.Println(err)
		return response, err
	}

	response = structs.DeleteEventResponse{Success: true}
	return response, err

}
