package api

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type ObservationAPIAdapter struct {
	query *db.Queries
}

func NewObservationAPIAdapter(query *db.Queries) *ObservationAPIAdapter {
	return &ObservationAPIAdapter{
		query: query,
	}
}

func (apiAdpt ObservationAPIAdapter) AddObservationApi(obsDateTime time.Time, bpSis, bpDia, pr, rr, temp float64, notes, patientId string, nurseId int64) (int64, error) {
	//theres no special core logic due to a this is a saving observation in database

	params := *&db.AddObservationParams{
		ObsDatetime: obsDateTime,
		BpSys:       bpSis,
		BpDia:       bpDia,
		Pr:          pr,
		Rr:          rr,
		Temp:        temp,
		Notes:       sql.NullString{String: notes, Valid: true},
		PatientID:   patientId,
		NurseID:     nurseId,
	}

	observation, err := apiAdpt.query.AddObservation(context.Background(), params)

	if err != nil {
		fmt.Println("AddObservation api error : ", err)
		return 0, err
	}

	return observation, nil

}

func (apiAdpt ObservationAPIAdapter) GetObservationByIDApi(observationId int64) (int64, time.Time, float64, float64, float64, float64, float64, string, string, int64, error) {

	observation, err := apiAdpt.query.GetObservationById(context.Background(), observationId)

	if err != nil {
		fmt.Println("GetObservationById api error : ", err)
		return 0, time.Time{}, 0.0, 0.0, 0.0, 0.0, 0.0, "", "", 0, err
	}

	return observation.ObservationID, observation.ObsDatetime, observation.BpSys, observation.BpDia, observation.Pr, observation.Rr, observation.Temp, observation.Notes.String, observation.PatientID, observation.NurseID, nil
}
func (apiAdpt ObservationAPIAdapter) UpdateObservationApi(obseravationIdIn int64, obsDateTimeIn time.Time, bpSysIn, bpDiaIn, prIn, rrIn, tempIn float64, notesIn, patientIdIn string, nurseIdIn int64) (int64, time.Time, float64, float64, float64, float64, float64, string, string, int64, error) {

	params := *&db.UpdateObservationParams{

		ObservationID: obseravationIdIn,
		ObsDatetime:   obsDateTimeIn,
		BpSys:         bpSysIn,
		BpDia:         bpDiaIn,
		Pr:            prIn,
		Rr:            rrIn,
		Temp:          tempIn,
		Notes:         sql.NullString{String: notesIn, Valid: true},
		PatientID:     patientIdIn,
		NurseID:       nurseIdIn,
	}
	observation, err := apiAdpt.query.UpdateObservation(context.Background(), params)

	if err != nil {
		fmt.Println("UpdateObservation api error : ", err)
		fmt.Println(err.Error())
		return 0, time.Time{}, 0.0, 0.0, 0.0, 0.0, 0.0, "", "", 0, err
	}

	return obseravationIdIn, observation.ObsDatetime, observation.BpSys, observation.BpDia, observation.Pr, observation.Pr, observation.Temp, observation.Notes.String, observation.PatientID, observation.NurseID, nil

}

//Get all blood pressure values in API layer					-Dhanuka Manuwansha						-25/05/2022
func (apiAdpt ObservationAPIAdapter) GetAllBPSisValuesApi(patientId string) ([]float64, error) {

	bpValues, err := apiAdpt.query.GetAllBPSisValues(context.Background(), patientId)

	if err != nil {
		fmt.Println("GetAllBPValues api error : ", err)
		var errorArray []float64
		return errorArray, err
	}

	return bpValues, nil
}

func (apiAdpt ObservationAPIAdapter) GetAllBPDiaValuesApi(patientId string) ([]float64, error) {

	bpValues, err := apiAdpt.query.GetAllBPDiaValues(context.Background(), patientId)

	if err != nil {
		fmt.Println("GetAllBPValues api error : ", err)
		var errorArray []float64
		return errorArray, err
	}

	return bpValues, nil
}

//Get all blood pressure values in API layer					-Dhanuka Manuwansha						-26/05/2022
func (apiAdpt ObservationAPIAdapter) GetAllPRValuesApi(patientId string) ([]float64, error) {

	prValues, err := apiAdpt.query.GetAllPRValues(context.Background(), patientId)

	if err != nil {
		fmt.Println("GetAllPRValues api error : ", err)
		var errorArray []float64
		return errorArray, err
	}

	return prValues, nil
}

//Get all respiratory rate values in API layer					-Dhanuka Manuwansha						-25/05/2022
func (apiAdpt ObservationAPIAdapter) GetAllRRValuesApi(patientId string) ([]float64, error) {

	rrValues, err := apiAdpt.query.GetAllRRValues(context.Background(), patientId)

	if err != nil {
		fmt.Println("GetAllRRValues api error : ", err)
		var errorArray []float64
		return errorArray, err
	}

	return rrValues, nil
}

// Get all temperature values in API layer					-Dhanuka Manuwansha						-25/05/2022
func (apiAdpt ObservationAPIAdapter) GetAllTempValuesApi(patientId string) ([]float64, error) {

	tempValues, err := apiAdpt.query.GetAllTempValues(context.Background(), patientId)

	if err != nil {
		fmt.Println("GetAllTempValues api error : ", err)
		var errorArray []float64
		return errorArray, err
	}

	return tempValues, nil
}

// Get all observations in API layer					-Dhanuka Manuwansha						-25/05/2022
func (apiAdpt ObservationAPIAdapter) GetAllObservationsApi(patientId string) ([]db.Observation, error) {

	observations, err := apiAdpt.query.GetAllObservations(context.Background(), patientId)

	if err != nil {
		fmt.Println("GetAllObservations api error : ", err)
		var errorArray []db.Observation
		return errorArray, err
	}

	return observations, nil
}
