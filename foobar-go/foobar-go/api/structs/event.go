package structs

type CreateEventRequest struct {
	Name       string `json:"name" binding:"required,min=3"`
	LocationId int64  `json:"location_id" binding:"gte=0"`
	Datetime   int64  `json:"datetime" binding:"required,gte=0"`
}

type CreateEventResponse struct {
	Success bool `json:"success"`
}

type GetEventByIdRequest struct {
	EventId int64 `uri:"id" binding:"gte=0"`
}

type GetEventByIdResponse struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Datetime int64  `json:"datetime"`
}
