package structs

type CreateEventRequest struct {
	Name       string `json:"name" binding:"required,min=3"`
	LocationId int64  `json:"location_id" binding:"gte=0"`
	Datetime   int64  `json:"datetime" binding:"required,gte=0"` //TODO: validation unix time
}

type CreateEventResponse struct {
	Success bool `json:"success"`
}

type GetEventByIdParamRequest struct {
	EventId int64 `uri:"id" binding:"gte=0"`
}

type GetEventByIdResponse struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Datetime int64  `json:"datetime"`
}

type GetEventListByPageQueryRequest struct {
	PageNumber int32 `form:"page_number" binding:"gte=0"`
	PageSize   int32 `form:"page_size" binding:"gte=0,lte=10"`
}

type EventJson struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Datetime int64  `json:"datetime"`
}

type GetEventListByPageResponse struct {
	Events []EventJson `json:"events"`
}

type UpdateEventParamRequest struct {
	EventId int64 `uri:"id" binding:"gte=0"`
}

type UpdateEventRequest struct {
	Name       string `json:"name" binding:"required,min=3"`
	LocationId int64  `json:"location_id" binding:"gte=0"`
	Datetime   int64  `json:"datetime" binding:"required,gte=0"` //TODO: validation unix time
}

type UpdateEventResponse struct {
	Success bool `json:"success"`
}
