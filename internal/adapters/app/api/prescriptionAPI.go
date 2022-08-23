package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

type PrescriptionAPIAdapter struct {
	query *db.Queries
}

func NewPrescriptionAPIAdapter(query *db.Queries) *PrescriptionAPIAdapter {
	return &PrescriptionAPIAdapter{
		query: query,
	}
}

func (apiAdpt PrescriptionAPIAdapter) GetActivePrescriptionsApi(patientId string) ([]db.Prescription, error) {

	prescriptions, err := apiAdpt.query.GetActivePrescriptions(context.Background(), patientId)

	if err != nil {
		fmt.Println("GetActivePrescription api error : ", err)
		var errorArray []db.Prescription
		return errorArray, err
	}

	return prescriptions, nil
}

func (apiAdpt PrescriptionAPIAdapter) GetLatestPrescriptionsApi(patientId string) ([]db.Prescription, error) {

	prescriptions, err := apiAdpt.query.GetLatestPrescriptions(context.Background(), patientId)

	if err != nil {
		fmt.Println("GetActivePrescription api error : ", err)
		var errorArray []db.Prescription
		return errorArray, err
	}

	return prescriptions, nil
}
