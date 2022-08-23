package ports

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
)

type DrugChartAPIPort interface {
	GetDrugChartApi(patientID string) ([]db.GetDrugChartRow, error)
	GetDrugchartForNurseDeskApi(arg db.GetDrugchartForNurseDeskParams) ([]db.GetDrugchartForNurseDeskRow, error)
}
