package service

import (
	"context"
	"foobar-go/api/structs"
	"foobar-go/infrastructure"
	db "foobar-go/sqlc"
	"foobar-go/util"
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

func (s EventService) CreateEvent(ctx context.Context, request structs.CreateEventRequest) (createEventResponse structs.CreateEventResponse, err error) {
	_, err = s.store.Queries.CreateEvent(ctx, db.CreateEventParams{
		Name:       request.Name,
		LocationID: request.LocationId,
		Datetime:   request.Datetime,
	})

	if err != nil {
		createEventResponse = structs.CreateEventResponse{
			Success: false,
		}

		if pgErr, isPGErr := err.(*pq.Error); isPGErr {
			if pgErr.Code == util.SQL_FOREIGN_KEY_VIOLATION_ERROR_CODE {
				return createEventResponse, &util.InvalidLocationId{}
			}
		}

		log.Println("error while creating event")
		log.Println(err)
		return createEventResponse, err

	}

	createEventResponse = structs.CreateEventResponse{
		Success: true,
	}

	return createEventResponse, err

}
