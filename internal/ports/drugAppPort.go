package ports

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
)

type DrugAPIPort interface {
	GetDrugByIdApi(drugID int64) (db.Drug, error)
}
