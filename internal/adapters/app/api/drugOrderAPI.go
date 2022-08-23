package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

type DrugOrderAPIAdapter struct {
	query *db.Queries
}

func NewDrugOrderAPIAdapter(query *db.Queries) *DrugOrderAPIAdapter {
	return &DrugOrderAPIAdapter{
		query: query,
	}
}

func (apiAdpt DrugOrderAPIAdapter) GetDrugOrdersByPrescriptionIDApi(prescriptionID int64) ([]db.DrugOrder, error) {

	drugOrders, err := apiAdpt.query.GetDrugOrdersByPrescriptionID(context.Background(), prescriptionID)
	fmt.Println("oooooooooooo")
	fmt.Println(drugOrders)
	if err != nil {
		fmt.Println("GetActivePrescription api error : ", err)
		var errorArray []db.DrugOrder
		return errorArray, err
	}

	return drugOrders, nil
}
