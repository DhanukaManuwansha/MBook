package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

type DrugChartAPIAdapter struct {
	query *db.Queries
}

func NewDrugChartAPIAdapter(query *db.Queries) *DrugChartAPIAdapter {
	return &DrugChartAPIAdapter{
		query: query,
	}
}

func (apiAdpt DrugChartAPIAdapter) GetDrugChartApi(patientId string) ([]db.GetDrugChartRow, error) {

	drug, err := apiAdpt.query.GetDrugChart(context.Background(), patientId)
	fmt.Println(drug)
	if err != nil {
		fmt.Println("GetDrugchart api error : ", err)
		var errorArray []db.GetDrugChartRow
		return errorArray, err
	}

	return drug, nil
}

func (apiAdpt DrugChartAPIAdapter) GetDrugchartForNurseDeskApi(arg db.GetDrugchartForNurseDeskParams) ([]db.GetDrugchartForNurseDeskRow, error) {

	drug, err := apiAdpt.query.GetDrugchartForNurseDesk(context.Background(), arg)
	fmt.Println(drug)
	if err != nil {
		fmt.Println("GetDrugchart api error : ", err)
		var errorArray []db.GetDrugchartForNurseDeskRow
		return errorArray, err
	}

	return drug, nil
}
