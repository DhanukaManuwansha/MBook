package ports

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
)

type CenterAPIPort interface {
	GetAllCenterIds() ([]db.Center, error)
}
