package ports

import (
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"time"
)

type NurseNoteAPIPort interface {
	UpdateNurseNoteApi(nursingNoteIdIn int64, remarkIn, nursingCareIn string, noteDateTimeIn time.Time, nurseIdIn int64, patientIdIn string) (int64, string, string, time.Time, int64, string, error)
	GetAllNurseNotesApi(patientId string) ([]db.NurseNotes, error)
	AddNurseNoteApi(remark, nursingCare string, notesDateTime string, createdAt string, nurseId string, patientId string) (int64, error)
	GetLatestNurseNotesApi(patientID string) ([]db.GetLatestNurseNotesRow, error)
}
