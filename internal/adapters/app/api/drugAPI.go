package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

type DrugAPIAdapter struct {
	query *db.Queries
}

func NewDrugAPIAdapter(query *db.Queries) *DrugAPIAdapter {
	return &DrugAPIAdapter{
		query: query,
	}
}

func (apiAdpt DrugAPIAdapter) GetDrugByIdApi(drugId int64) (db.Drug, error) {

	drug, err := apiAdpt.query.GetDrugById(context.Background(), drugId)

	if err != nil {
		fmt.Println("GetActivePrescription api error : ", err)
		var errorArray db.Drug
		return errorArray, err
	}

	return drug, nil
}
