package ports

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
)

type DrugOrderAPIPort interface {
	GetDrugOrdersByPrescriptionIDApi(prescriptionID int64) ([]db.DrugOrder, error)
}
