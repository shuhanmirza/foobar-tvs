package controller

import (
	"foobar-go/api/service"
	"foobar-go/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConfigurationController struct {
	configurationService service.ConfigurationService
}

func NewConfigurationController(configurationService service.ConfigurationService) ConfigurationController {
	return ConfigurationController{
		configurationService: configurationService,
	}
}

func (c ConfigurationController) GetLocationList(ctx *gin.Context) {
	getLocationListResponse, err := c.configurationService.GetLocationList(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, getLocationListResponse)
}
