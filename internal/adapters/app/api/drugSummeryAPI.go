package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

type DrugSummeryAPIAdapter struct {
	query *db.Queries
}

func NewDrugSummeryAPIAdapter(query *db.Queries) *DrugSummeryAPIAdapter {
	return &DrugSummeryAPIAdapter{
		query: query,
	}
}

func (apiAdpt DrugSummeryAPIAdapter) GetAllDrugSummeryApi(patientId string) ([]db.DrugSummery, error) {

	drug, err := apiAdpt.query.GetAllDrugSummery(context.Background(), patientId)

	if err != nil {
		fmt.Println("GetActivePrescription api error : ", err)
		var errorArray []db.DrugSummery
		return errorArray, err
	}

	return drug, nil
}

func (apiAdpt DrugSummeryAPIAdapter) GetAllDrugSummeryOfADrugApi(arg db.GetAllDrugSummeryOfADrugParams) ([]db.GetAllDrugSummeryOfADrugRow, error) {

	drug, err := apiAdpt.query.GetAllDrugSummeryOfADrug(context.Background(), arg)

	if err != nil {
		fmt.Println("GetActivePrescription api error : ", err)
		var errorArray []db.GetAllDrugSummeryOfADrugRow
		return errorArray, err
	}

	return drug, nil
}

func (apiAdpt DrugSummeryAPIAdapter) UpdateDrugSummeryApi(arg db.UpdateDrugSummeryParams) (db.DrugSummery, error) {

	summery, err := apiAdpt.query.UpdateDrugSummery(context.Background(), arg)

	if err != nil {
		fmt.Println("UpdateObservation api error : ", err)
		fmt.Println(err.Error())
		var errorObj db.DrugSummery
		return errorObj, err
	}

	return summery, nil

}
