package api

import (
	patientdb "MedbookServer/internal/adapters/framework/driven/db/sqlc/patientDB_sqlc"
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

type PatientAPIAdapter struct {
	query *patientdb.Queries
}

func NewPatientAPIAdapter(query *patientdb.Queries) *PatientAPIAdapter {
	return &PatientAPIAdapter{
		query: query,
	}
}

func (apiAdpt PatientAPIAdapter) GetPatientsApi() ([]patientdb.Patient, error) {

	patients, err := apiAdpt.query.GetPatients(context.Background())

	if err != nil {
		fmt.Println("GetActivePrescription api error : ", err)
		var errorArray []patientdb.Patient
		return errorArray, err
	}

	return patients, nil
}

func (apiAdpt PatientAPIAdapter) GetPatientBYNICApi(nic string) ([]patientdb.Patient, error) {

	patients, err := apiAdpt.query.GetPatientBYNIC(context.Background(), nic)

	if err != nil {
		fmt.Println("GetActivePrescription api error : ", err)
		var errorArray []patientdb.Patient
		return errorArray, err
	}

	return patients, nil
}

func (apiAdpt PatientAPIAdapter) GetPatientBYNameApi(name string) ([]patientdb.Patient, error) {

	patients, err := apiAdpt.query.GetPatientBYName(context.Background(), name)

	if err != nil {
		fmt.Println("GetActivePrescription api error : ", err)
		var errorArray []patientdb.Patient
		return errorArray, err
	}

	return patients, nil
}

func (apiAdpt PatientAPIAdapter) GetPatientBYTellNoApi(tellNo string) ([]patientdb.Patient, error) {

	patients, err := apiAdpt.query.GetPatientBYTellNo(context.Background(), tellNo)

	if err != nil {
		fmt.Println("GetActivePrescription api error : ", err)
		var errorArray []patientdb.Patient
		return errorArray, err
	}

	return patients, nil
}

func (apiAdpt PatientAPIAdapter) GetPatientByIdApi(patientId string) (patientdb.Patient, error) {

	patient, err := apiAdpt.query.GetPatientById(context.Background(), patientId)

	if err != nil {
		fmt.Println("GetActivePrescription api error : ", err)
		var errorObject patientdb.Patient
		return errorObject, err
	}

	return patient, nil
}
