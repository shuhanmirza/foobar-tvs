package controller

import (
	"foobar-go/api/service"
	"foobar-go/api/structs"
	"foobar-go/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EventController struct {
	eventService service.EventService
}

func NewEventController(eventService service.EventService) EventController {
	return EventController{
		eventService: eventService,
	}
}

func (c EventController) CreateEvent(ctx *gin.Context) {
	var request structs.CreateEventRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		util.HandleCommonValidationError(ctx, err)
		return
	}

	createEventResponse, err := c.eventService.CreateEvent(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, createEventResponse)
}

func (c EventController) GetEventById(ctx *gin.Context) {
	var request structs.GetEventByIdParamRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		util.HandleCommonValidationError(ctx, err)
		return
	}

	response, err := c.eventService.GetEventById(ctx, request.EventId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c EventController) GetEventListByPage(ctx *gin.Context) {
	var request structs.GetEventListByPageQueryRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		util.HandleCommonValidationError(ctx, err)
		return
	}

	response, err := c.eventService.GetEventListByPage(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c EventController) UpdateEvent(ctx *gin.Context) {
	var requestParam structs.UpdateEventParamRequest
	var request structs.UpdateEventRequest

	if err := ctx.ShouldBindUri(&requestParam); err != nil {
		util.HandleCommonValidationError(ctx, err)
		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		util.HandleCommonValidationError(ctx, err)
		return
	}

	response, err := c.eventService.UpdateEvent(ctx, requestParam.EventId, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c EventController) DeleteEvent(ctx *gin.Context) {
	var request structs.DeleteEventParamRequest

	if err := ctx.ShouldBindUri(&request); err != nil {
		util.HandleCommonValidationError(ctx, err)
		return
	}

	response, err := c.eventService.DeleteEvent(ctx, request.EventId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
