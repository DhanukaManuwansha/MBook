package ports

import (
	patientdb "MedbookServer/internal/adapters/framework/driven/db/sqlc/patientDB_sqlc"
)

type PatientAPIPort interface {
	GetPatientsApi() ([]patientdb.Patient, error)
	GetPatientBYNICApi(nic string) ([]patientdb.Patient, error)
	GetPatientBYNameApi(firstName string) ([]patientdb.Patient, error)
	GetPatientBYTellNoApi(tellNo string) ([]patientdb.Patient, error)
	GetPatientByIdApi(patientID string) (patientdb.Patient, error)
}
