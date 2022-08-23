package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
)

type DrugSummeryAPIPort interface {
	GetAllDrugSummeryApi(cpatientID string) ([]db.DrugSummery, error)
	GetAllDrugSummeryOfADrugApi(arg db.GetAllDrugSummeryOfADrugParams) ([]db.GetAllDrugSummeryOfADrugRow, error)
	UpdateDrugSummeryApi(arg db.UpdateDrugSummeryParams) (db.DrugSummery, error)
}
