package ports

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"time"
)

type ObservationAPIPort interface {
	AddObservationApi(obsDateTime time.Time, bpSis, bpDia, pr, rr, temp float64, notes, patientId string, nurseId int64) (int64, error)
	GetObservationByIDApi(observationId int64) (int64, time.Time, float64, float64, float64, float64, float64, string, string, int64, error)
	UpdateObservationApi(obseravationIdIn int64, obsDateTimeIn time.Time, bpSisIn, bpDiaIn, prIn, rrIn, tempIn float64, notesIn, patientIdIn string, nurseIdIn int64) (int64, time.Time, float64, float64, float64, float64, float64, string, string, int64, error)
	GetAllBPSisValuesApi(patientID string) ([]float64, error)
	GetAllBPDiaValuesApi(patientID string) ([]float64, error)
	GetAllPRValuesApi(patientId string) ([]float64, error)
	GetAllRRValuesApi(patientId string) ([]float64, error)
	GetAllTempValuesApi(patientId string) ([]float64, error)
	GetAllObservationsApi(patientId string) ([]db.Observation, error)
	
}
