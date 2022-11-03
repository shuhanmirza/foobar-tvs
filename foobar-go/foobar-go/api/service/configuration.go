package service

import (
	"foobar-go/api/structs"
	"foobar-go/infrastructure"
	"github.com/gin-gonic/gin"
	"log"
)

type ConfigurationService struct {
	store *infrastructure.Store
}

func NewConfigurationService(store *infrastructure.Store) ConfigurationService {
	return ConfigurationService{
		store: store,
	}
}

func (s ConfigurationService) GetLocationList(ctx *gin.Context) (locationListResponse structs.GetLocationListResponse, err error) {
	locationList, err := s.store.Queries.ListLocations(ctx)
	if err != nil {
		log.Println("error while retrieving location list")
		log.Println(err)
		return locationListResponse, err
	}

	locationListResponse = structs.GetLocationListResponse{
		LocationList: locationList,
	}

	return locationListResponse, err

}
