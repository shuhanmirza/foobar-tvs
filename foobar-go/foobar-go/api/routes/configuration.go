package routes

import (
	"foobar-go/api/controller"
	"foobar-go/infrastructure"
)

type ConfigurationRoute struct {
	Controller controller.ConfigurationController
	Handler    infrastructure.GinRouter
}

func NewConfigurationRoute(configurationController controller.ConfigurationController,
	ginRouter infrastructure.GinRouter) ConfigurationRoute {
	return ConfigurationRoute{
		Controller: configurationController,
		Handler:    ginRouter,
	}
}

func (r ConfigurationRoute) Setup() {
	configuration := r.Handler.Gin.Group("config")
	{
		configuration.GET("/location-list", r.Controller.GetLocationList)
	}
}
