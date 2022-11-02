package routes

import (
	"foobar-go/api/controller"
	"foobar-go/infrastructure"
)

type EventRoute struct {
	Controller controller.EventController
	Handler    infrastructure.GinRouter
}

func NewEventRoute(eventController controller.EventController,
	ginRouter infrastructure.GinRouter) EventRoute {
	return EventRoute{
		Controller: eventController,
		Handler:    ginRouter,
	}
}

func (r EventRoute) Setup() {
	configuration := r.Handler.Gin.Group("api/events")
	{
		configuration.POST("/", r.Controller.CreateEvent)
	}
}
