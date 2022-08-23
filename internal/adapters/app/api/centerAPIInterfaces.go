package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
)

type CenterAPI interface {
	GetAllCenterIds() ([]db.Center, error)
}
