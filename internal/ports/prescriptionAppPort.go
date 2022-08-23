package ports

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
)

type PrescriptionAPIPort interface {
	GetActivePrescriptionsApi(patientID string) ([]db.Prescription, error)
	GetLatestPrescriptionsApi(patientID string) ([]db.Prescription, error)
}
