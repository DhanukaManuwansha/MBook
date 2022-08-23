package ports

import db "MedbookServer/internal/adapters/framework/driven/db/sqlc"

type BedAPIPort interface {
	// GetBedsbyWard
	GetBedsByWardApi(wardID int64) ([]db.GetBedsByWardRow, error)
}
