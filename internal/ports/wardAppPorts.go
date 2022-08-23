package ports

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
)

type WardAPIPort interface {
	//GetWardsByCenterAPI
	GetWardsByCenterAPI(center_id int64) ([]db.Ward, error)
}
