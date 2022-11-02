package structs

import db "foobar-go/sqlc"

type GetLocationListResponse struct {
	LocationList []db.Locations `json:"location_list"`
}
