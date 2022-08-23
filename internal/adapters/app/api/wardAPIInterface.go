package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
)

type wardAPI interface {
	GetWardsByCenterAPI(center_id int64) ([]db.Ward, error)
}
